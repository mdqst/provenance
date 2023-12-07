package app

// DONTCOVER

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"runtime/debug"
	"sort"
	"strings"
	"testing"
	"time"

	icqtypes "github.com/cosmos/ibc-apps/modules/async-icq/v6/types"
	"github.com/stretchr/testify/require"

	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdksim "github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/cosmos/cosmos-sdk/x/quarantine"
	"github.com/cosmos/cosmos-sdk/x/sanction"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	icagenesistypes "github.com/cosmos/ibc-go/v6/modules/apps/27-interchain-accounts/genesis/types"
	icatypes "github.com/cosmos/ibc-go/v6/modules/apps/27-interchain-accounts/types"

	cmdconfig "github.com/provenance-io/provenance/cmd/provenanced/config"
	"github.com/provenance-io/provenance/internal/pioconfig"
	attributetypes "github.com/provenance-io/provenance/x/attribute/types"
	"github.com/provenance-io/provenance/x/hold"
	markertypes "github.com/provenance-io/provenance/x/marker/types"
	metadatatypes "github.com/provenance-io/provenance/x/metadata/types"
	msgfeetype "github.com/provenance-io/provenance/x/msgfees/types"
	nametypes "github.com/provenance-io/provenance/x/name/types"
	triggertypes "github.com/provenance-io/provenance/x/trigger/types"
)

func init() {
	sdksim.GetSimulatorFlags()
	pioconfig.SetProvenanceConfig("", 0)
}

type StoreKeysPrefixes struct {
	A        storetypes.StoreKey
	B        storetypes.StoreKey
	Prefixes [][]byte
}

// ProvAppStateFn wraps the sdksim.AppStateFn and sets the ICA GenesisState if isn't yet defined in the appState.
func ProvAppStateFn(cdc codec.JSONCodec, simManager *module.SimulationManager) simtypes.AppStateFn {
	return func(r *rand.Rand, accs []simtypes.Account, config simtypes.Config) (json.RawMessage, []simtypes.Account, string, time.Time) {
		appState, simAccs, chainID, genesisTimestamp := sdksim.AppStateFn(cdc, simManager)(r, accs, config)
		appState = appStateWithICA(appState, cdc)
		appState = appStateWithICQ(appState, cdc)
		return appState, simAccs, chainID, genesisTimestamp
	}
}

// appStateWithICA checks the given appState for an ica entry. If it's not found, it's populated with the defaults.
func appStateWithICA(appState json.RawMessage, cdc codec.JSONCodec) json.RawMessage {
	rawState := make(map[string]json.RawMessage)
	err := json.Unmarshal(appState, &rawState)
	if err != nil {
		panic(fmt.Sprintf("error unmarshalling appstate: %v", err))
	}
	icaGenJSON, icaGenFound := rawState[icatypes.ModuleName]
	if !icaGenFound || len(icaGenJSON) == 0 {
		icaGenState := icagenesistypes.DefaultGenesis()
		rawState[icatypes.ModuleName] = cdc.MustMarshalJSON(icaGenState)
		appState, err = json.Marshal(rawState)
		if err != nil {
			panic(fmt.Sprintf("error marshalling appstate: %v", err))
		}
	}
	return appState
}

// appStateWithICA checks the given appState for an ica entry. If it's not found, it's populated with the defaults.
func appStateWithICQ(appState json.RawMessage, cdc codec.JSONCodec) json.RawMessage {
	rawState := make(map[string]json.RawMessage)
	err := json.Unmarshal(appState, &rawState)
	if err != nil {
		panic(fmt.Sprintf("error unmarshalling appstate: %v", err))
	}
	icqGenJSON, icqGenFound := rawState[icqtypes.ModuleName]
	if !icqGenFound || len(icqGenJSON) == 0 {
		icqGenState := icqtypes.DefaultGenesis()
		rawState[icqtypes.ModuleName] = cdc.MustMarshalJSON(icqGenState)
		appState, err = json.Marshal(rawState)
		if err != nil {
			panic(fmt.Sprintf("error marshalling appstate: %v", err))
		}
	}
	return appState
}

func TestFullAppSimulation(t *testing.T) {
	config, db, dir, logger, skip, err := sdksim.SetupSimulation("leveldb-app-sim", "Simulation")
	if skip {
		t.Skip("skipping provenance application simulation")
	}
	PrintConfig(config)
	require.NoError(t, err, "provenance simulation setup failed")

	defer func() {
		db.Close()
		require.NoError(t, os.RemoveAll(dir))
	}()

	app := New(logger, db, nil, true, map[int64]bool{}, t.TempDir(), sdksim.FlagPeriodValue, MakeEncodingConfig(), sdksim.EmptyAppOptions{}, fauxMerkleModeOpt)
	require.Equal(t, "provenanced", app.Name())

	fmt.Printf("running provenance full app simulation\n")

	// run randomized simulation
	_, _, simParams, simErr := simulation.SimulateFromSeed(
		t,
		os.Stdout,
		app.BaseApp,
		ProvAppStateFn(app.AppCodec(), app.SimulationManager()),
		simtypes.RandomAccounts, // Replace with own random account function if using keys other than secp256k1
		sdksim.SimulationOperations(app, app.AppCodec(), config),
		app.ModuleAccountAddrs(),
		config,
		app.AppCodec(),
	)

	// export state and simParams before the simulation error is checked
	err = sdksim.CheckExportSimulation(app, config, simParams)
	require.NoError(t, err)
	require.NoError(t, simErr)

	PrintStats(config, db)
}

func TestSimple(t *testing.T) {
	config, db, dir, logger, skip, err := sdksim.SetupSimulation("leveldb-app-sim", "Simulation")
	if skip {
		t.Skip("skipping provenance application simulation")
	}
	PrintConfig(config)
	require.NoError(t, err, "provenance simulation setup failed")

	defer func() {
		db.Close()
		require.NoError(t, os.RemoveAll(dir))
	}()

	app := New(logger, db, nil, true, map[int64]bool{}, t.TempDir(), sdksim.FlagPeriodValue, MakeEncodingConfig(), sdksim.EmptyAppOptions{}, fauxMerkleModeOpt)
	require.Equal(t, "provenanced", app.Name())

	// run randomized simulation
	_, _, _, simErr := simulation.SimulateFromSeed(
		t,
		os.Stdout,
		app.BaseApp,
		ProvAppStateFn(app.AppCodec(), app.SimulationManager()),
		simtypes.RandomAccounts, // Replace with own random account function if using keys other than secp256k1
		sdksim.SimulationOperations(app, app.AppCodec(), config),
		app.ModuleAccountAddrs(),
		config,
		app.AppCodec(),
	)

	require.NoError(t, simErr)
	PrintStats(config, db)
}

// Profile with:
// /usr/local/go/bin/go test -benchmem -run=^$ github.com/provenance-io/provenance -bench ^BenchmarkFullAppSimulation$ -Commit=true -cpuprofile cpu.out
func TestAppImportExport(t *testing.T) {
	// uncomment to run in ide without flags.
	//sdksim.FlagEnabledValue = true

	config, db, dir, logger, skip, err := sdksim.SetupSimulation("leveldb-app-sim", "Simulation")
	if skip {
		t.Skip("skipping application import/export simulation")
	}
	PrintConfig(config)
	require.NoError(t, err, "simulation setup failed")

	defer func() {
		require.NoError(t, db.Close())
		require.NoError(t, os.RemoveAll(dir))
	}()

	home := t.TempDir()
	app := New(logger, db, nil, true, map[int64]bool{}, home, sdksim.FlagPeriodValue, MakeEncodingConfig(), sdksim.EmptyAppOptions{}, fauxMerkleModeOpt)

	fmt.Printf("running provenance test import export\n")

	// Run randomized simulation
	_, lastBlockTime, simParams, simErr := simulation.SimulateFromSeed(
		t,
		os.Stdout,
		app.BaseApp,
		ProvAppStateFn(app.AppCodec(), app.SimulationManager()),
		simtypes.RandomAccounts, // Replace with own random account function if using keys other than secp256k1
		sdksim.SimulationOperations(app, app.AppCodec(), config),
		app.ModuleAccountAddrs(),
		config,
		app.AppCodec(),
	)

	// export state and simParams before the simulation error is checked
	err = sdksim.CheckExportSimulation(app, config, simParams)
	require.NoError(t, err)
	require.NoError(t, simErr)

	PrintStats(config, db)

	fmt.Printf("exporting genesis...\n")

	exported, err := app.ExportAppStateAndValidators(false, []string{})
	require.NoError(t, err)

	fmt.Printf("importing genesis...\n")

	_, newDB, newDir, _, _, err := sdksim.SetupSimulation("leveldb-app-sim-2", "Simulation-2")
	require.NoError(t, err, "simulation setup failed")

	defer func() {
		require.NoError(t, newDB.Close())
		require.NoError(t, os.RemoveAll(newDir))
	}()

	newApp := New(log.NewNopLogger(), newDB, nil, true, map[int64]bool{}, home, sdksim.FlagPeriodValue, MakeEncodingConfig(), sdksim.EmptyAppOptions{}, fauxMerkleModeOpt)

	var genesisState sdksim.GenesisState
	err = json.Unmarshal(exported.AppState, &genesisState)
	require.NoError(t, err)

	defer func() {
		if r := recover(); r != nil {
			rstr := fmt.Sprintf("%v", r)
			if !strings.Contains(rstr, "validator set is empty after InitGenesis") {
				panic(r)
			}
			logger.Info("Skipping simulation as all validators have been unbonded")
			logger.Info("err", rstr, "stacktrace", string(debug.Stack()))
		}
	}()

	ctxA := app.NewContext(true, tmproto.Header{Height: app.LastBlockHeight(), Time: lastBlockTime})
	ctxB := newApp.NewContext(true, tmproto.Header{Height: app.LastBlockHeight(), Time: lastBlockTime})
	newApp.mm.InitGenesis(ctxB, app.AppCodec(), genesisState)
	newApp.StoreConsensusParams(ctxB, exported.ConsensusParams)

	fmt.Printf("comparing stores...\n")

	storeKeysPrefixes := []StoreKeysPrefixes{
		{app.GetKey(authtypes.StoreKey), newApp.GetKey(authtypes.StoreKey), [][]byte{}},
		{app.GetKey(stakingtypes.StoreKey), newApp.GetKey(stakingtypes.StoreKey),
			[][]byte{
				stakingtypes.UnbondingQueueKey, stakingtypes.RedelegationQueueKey, stakingtypes.ValidatorQueueKey,
				stakingtypes.HistoricalInfoKey,
			}}, // ordering may change but it doesn't matter
		{app.GetKey(slashingtypes.StoreKey), newApp.GetKey(slashingtypes.StoreKey), [][]byte{}},
		{app.GetKey(minttypes.StoreKey), newApp.GetKey(minttypes.StoreKey), [][]byte{}},
		{app.GetKey(distrtypes.StoreKey), newApp.GetKey(distrtypes.StoreKey), [][]byte{}},
		{app.GetKey(banktypes.StoreKey), newApp.GetKey(banktypes.StoreKey), [][]byte{banktypes.BalancesPrefix}},
		{app.GetKey(paramtypes.StoreKey), newApp.GetKey(paramtypes.StoreKey), [][]byte{}},
		{app.GetKey(govtypes.StoreKey), newApp.GetKey(govtypes.StoreKey), [][]byte{}},
		{app.GetKey(evidencetypes.StoreKey), newApp.GetKey(evidencetypes.StoreKey), [][]byte{}},
		{app.GetKey(capabilitytypes.StoreKey), newApp.GetKey(capabilitytypes.StoreKey), [][]byte{}},
		{app.GetKey(authzkeeper.StoreKey), newApp.GetKey(authzkeeper.StoreKey), [][]byte{authzkeeper.GrantKey, authzkeeper.GrantQueuePrefix}},
		{app.GetKey(quarantine.StoreKey), newApp.GetKey(quarantine.StoreKey), [][]byte{}},
		{app.GetKey(sanction.StoreKey), newApp.GetKey(sanction.StoreKey), [][]byte{}},

		{app.GetKey(markertypes.StoreKey), newApp.GetKey(markertypes.StoreKey), [][]byte{}},
		{app.GetKey(msgfeetype.StoreKey), newApp.GetKey(msgfeetype.StoreKey), [][]byte{}},
		{app.GetKey(attributetypes.StoreKey), newApp.GetKey(attributetypes.StoreKey), [][]byte{attributetypes.AttributeAddrLookupKeyPrefix}},
		{app.GetKey(nametypes.StoreKey), newApp.GetKey(nametypes.StoreKey), [][]byte{}},
		{app.GetKey(metadatatypes.StoreKey), newApp.GetKey(metadatatypes.StoreKey), [][]byte{}},
		{app.GetKey(triggertypes.StoreKey), newApp.GetKey(triggertypes.StoreKey), [][]byte{}},
		{app.GetKey(hold.StoreKey), newApp.GetKey(hold.StoreKey), [][]byte{}},
	}

	for _, skp := range storeKeysPrefixes {
		storeA := ctxA.KVStore(skp.A)
		storeB := ctxB.KVStore(skp.B)

		failedKVAs, failedKVBs := sdk.DiffKVStores(storeA, storeB, skp.Prefixes)
		require.Equal(t, len(failedKVAs), len(failedKVBs), "unequal sets of key-values to compare")

		fmt.Printf("compared %d different key/value pairs between %s and %s\n", len(failedKVAs), skp.A, skp.B)
		require.Equal(t, 0, len(failedKVAs), sdksim.GetSimulationLog(skp.A.Name(), app.SimulationManager().StoreDecoders, failedKVAs, failedKVBs))
	}
}

func TestAppSimulationAfterImport(t *testing.T) {
	config, db, dir, logger, skip, err := sdksim.SetupSimulation("leveldb-app-sim", "Simulation")
	if skip {
		t.Skip("skipping application simulation after import")
	}
	PrintConfig(config)
	require.NoError(t, err, "simulation setup failed")

	defer func() {
		db.Close()
		require.NoError(t, os.RemoveAll(dir))
	}()

	home := t.TempDir()
	app := New(logger, db, nil, true, map[int64]bool{}, home, sdksim.FlagPeriodValue, MakeEncodingConfig(), sdksim.EmptyAppOptions{}, fauxMerkleModeOpt)

	// Run randomized simulation
	stopEarly, lastBlockTime, simParams, simErr := simulation.SimulateFromSeed(
		t,
		os.Stdout,
		app.BaseApp,
		ProvAppStateFn(app.AppCodec(), app.SimulationManager()),
		simtypes.RandomAccounts, // Replace with own random account function if using keys other than secp256k1
		sdksim.SimulationOperations(app, app.AppCodec(), config),
		app.ModuleAccountAddrs(),
		config,
		app.AppCodec(),
	)

	// export state and simParams before the simulation error is checked
	err = sdksim.CheckExportSimulation(app, config, simParams)
	require.NoError(t, err)
	require.NoError(t, simErr)

	PrintStats(config, db)

	if stopEarly {
		fmt.Println("can't export or import a zero-validator genesis, exiting test...")
		return
	}

	fmt.Printf("exporting genesis...\n")

	exported, err := app.ExportAppStateAndValidators(true, []string{})
	require.NoError(t, err)

	fmt.Printf("importing genesis...\n")

	_, newDB, newDir, _, _, err := sdksim.SetupSimulation("leveldb-app-sim-2", "Simulation-2")
	require.NoError(t, err, "simulation setup failed")

	defer func() {
		newDB.Close()
		require.NoError(t, os.RemoveAll(newDir))
	}()

	newApp := New(log.NewNopLogger(), newDB, nil, true, map[int64]bool{}, home, sdksim.FlagPeriodValue, MakeEncodingConfig(), sdksim.EmptyAppOptions{}, fauxMerkleModeOpt)

	newApp.InitChain(abci.RequestInitChain{
		AppStateBytes: exported.AppState,
		Time:          lastBlockTime,
	})

	sdksim.FlagGenesisTimeValue = lastBlockTime.Unix()
	_, _, _, err = simulation.SimulateFromSeed(
		t,
		os.Stdout,
		newApp.BaseApp,
		ProvAppStateFn(app.AppCodec(), app.SimulationManager()),
		simtypes.RandomAccounts, // Replace with own random account function if using keys other than secp256k1
		sdksim.SimulationOperations(newApp, newApp.AppCodec(), config),
		app.ModuleAccountAddrs(),
		config,
		app.AppCodec(),
	)
	require.NoError(t, err)
}

// TODO: Make another test for the fuzzer itself, which just has noOp txs
// and doesn't depend on the application.
func TestAppStateDeterminism(t *testing.T) {
	// uncomment these to run in ide without flags.
	//sdksim.FlagEnabledValue = true
	//sdksim.FlagBlockSizeValue = 100
	//sdksim.FlagNumBlocksValue = 50

	if !sdksim.FlagEnabledValue {
		t.Skip("skipping application simulation")
	}

	config := sdksim.NewConfigFromFlags()
	config.InitialBlockHeight = 1
	config.ExportParamsPath = ""
	config.OnOperation = false
	config.AllInvariants = false
	config.ChainID = helpers.SimAppChainID
	config.DBBackend = "memdb"
	config.Commit = true

	numSeeds := 3
	numTimesToRunPerSeed := 5
	appHashList := make([]json.RawMessage, numTimesToRunPerSeed)

	home := t.TempDir()

	seeds := make([]int64, numSeeds)
	for i := range seeds {
		seeds[i] = rand.Int63()
	}

	// uncomment and tweak to use a single specific seed.
	//seeds = []int64{9171851189930047994}

	for i, seed := range seeds {
		config.Seed = seed
		PrintConfig(config)

		for j := 0; j < numTimesToRunPerSeed; j++ {
			var logger log.Logger
			if sdksim.FlagVerboseValue {
				logger = log.TestingLogger()
			} else {
				logger = log.NewNopLogger()
			}

			db := dbm.NewMemDB()
			app := New(logger, db, nil, true, map[int64]bool{}, home, sdksim.FlagPeriodValue, MakeEncodingConfig(), sdksim.EmptyAppOptions{}, interBlockCacheOpt())

			fmt.Printf(
				"running provenance non-determinism simulation; seed %d: %d/%d, attempt: %d/%d\n",
				config.Seed, i+1, numSeeds, j+1, numTimesToRunPerSeed,
			)

			_, _, _, err := simulation.SimulateFromSeed(
				t,
				os.Stdout,
				app.BaseApp,
				ProvAppStateFn(app.AppCodec(), app.SimulationManager()),
				simtypes.RandomAccounts, // Replace with own random account function if using keys other than secp256k1
				sdksim.SimulationOperations(app, app.AppCodec(), config),
				app.ModuleAccountAddrs(),
				config,
				app.AppCodec(),
			)
			require.NoError(t, err)

			PrintStats(config, db)

			appHash := app.LastCommitID().Hash
			appHashList[j] = appHash

			if j != 0 {
				require.Equal(
					t, string(appHashList[0]), string(appHashList[j]),
					"non-determinism in seed %d: %d/%d, attempt: %d/%d\n", config.Seed, i+1, numSeeds, j+1, numTimesToRunPerSeed,
				)
			}
		}
	}
}

// fauxMerkleModeOpt returns a BaseApp option to use a dbStoreAdapter instead of
// an IAVLStore for faster simulation speed.
func fauxMerkleModeOpt(bapp *baseapp.BaseApp) {
	bapp.SetFauxMerkleMode()
}

// interBlockCacheOpt returns a BaseApp option function that sets the persistent
// inter-block write-through cache.
func interBlockCacheOpt() func(*baseapp.BaseApp) {
	return baseapp.SetInterBlockCache(store.NewCommitKVStoreCacheManager())
}

// PrintStats outputs the config and db info.
func PrintStats(config simtypes.Config, db dbm.DB) {
	PrintConfig(config)
	if config.Commit {
		PrintDBInfo(db)
	}
}

// PrintConfig outputs the config.
func PrintConfig(config simtypes.Config) {
	fmt.Println("-vvv-  Config Info  -vvv-")
	cfields := cmdconfig.MakeFieldValueMap(config, true)
	for _, f := range cfields.GetSortedKeys() {
		fmt.Printf("%s: %s\n", f, cfields.GetStringOf(f))
	}
	fmt.Println("-^^^-  Config Info  -^^^-")
}

// PrintDBInfo outputs the db.Stats map.
func PrintDBInfo(db dbm.DB) {
	fmt.Println("-vvv-  Database Info  -vvv-")
	dbStats := db.Stats()
	if len(dbStats) == 0 {
		fmt.Println("No info to report.")
	} else {
		keys := make([]string, 0, len(dbStats))
		for k := range dbStats {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			v := dbStats[k]
			if strings.Contains(v, "\n") {
				fmt.Printf("%s:\n", k)
				fmt.Println(v)
			} else {
				fmt.Printf("%s: %q\n", k, v)
			}
		}
	}
	fmt.Println("-^^^-  Database Info  -^^^-")
}
