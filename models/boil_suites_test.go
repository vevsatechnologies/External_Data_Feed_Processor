// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("BittrexHistoricData", testBittrexHistoricData)
	t.Run("Data", testData)
	t.Run("PoloniexHistoricData", testPoloniexHistoricData)
	t.Run("Posdatatables", testPosdatatables)
	t.Run("Powdatatables", testPowdatatables)
}

func TestDelete(t *testing.T) {
	t.Run("BittrexHistoricData", testBittrexHistoricDataDelete)
	t.Run("Data", testDataDelete)
	t.Run("PoloniexHistoricData", testPoloniexHistoricDataDelete)
	t.Run("Posdatatables", testPosdatatablesDelete)
	t.Run("Powdatatables", testPowdatatablesDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("BittrexHistoricData", testBittrexHistoricDataQueryDeleteAll)
	t.Run("Data", testDataQueryDeleteAll)
	t.Run("PoloniexHistoricData", testPoloniexHistoricDataQueryDeleteAll)
	t.Run("Posdatatables", testPosdatatablesQueryDeleteAll)
	t.Run("Powdatatables", testPowdatatablesQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("BittrexHistoricData", testBittrexHistoricDataSliceDeleteAll)
	t.Run("Data", testDataSliceDeleteAll)
	t.Run("PoloniexHistoricData", testPoloniexHistoricDataSliceDeleteAll)
	t.Run("Posdatatables", testPosdatatablesSliceDeleteAll)
	t.Run("Powdatatables", testPowdatatablesSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("BittrexHistoricData", testBittrexHistoricDataExists)
	t.Run("Data", testDataExists)
	t.Run("PoloniexHistoricData", testPoloniexHistoricDataExists)
	t.Run("Posdatatables", testPosdatatablesExists)
	t.Run("Powdatatables", testPowdatatablesExists)
}

func TestFind(t *testing.T) {
	t.Run("BittrexHistoricData", testBittrexHistoricDataFind)
	t.Run("Data", testDataFind)
	t.Run("PoloniexHistoricData", testPoloniexHistoricDataFind)
	t.Run("Posdatatables", testPosdatatablesFind)
	t.Run("Powdatatables", testPowdatatablesFind)
}

func TestBind(t *testing.T) {
	t.Run("BittrexHistoricData", testBittrexHistoricDataBind)
	t.Run("Data", testDataBind)
	t.Run("PoloniexHistoricData", testPoloniexHistoricDataBind)
	t.Run("Posdatatables", testPosdatatablesBind)
	t.Run("Powdatatables", testPowdatatablesBind)
}

func TestOne(t *testing.T) {
	t.Run("BittrexHistoricData", testBittrexHistoricDataOne)
	t.Run("Data", testDataOne)
	t.Run("PoloniexHistoricData", testPoloniexHistoricDataOne)
	t.Run("Posdatatables", testPosdatatablesOne)
	t.Run("Powdatatables", testPowdatatablesOne)
}

func TestAll(t *testing.T) {
	t.Run("BittrexHistoricData", testBittrexHistoricDataAll)
	t.Run("Data", testDataAll)
	t.Run("PoloniexHistoricData", testPoloniexHistoricDataAll)
	t.Run("Posdatatables", testPosdatatablesAll)
	t.Run("Powdatatables", testPowdatatablesAll)
}

func TestCount(t *testing.T) {
	t.Run("BittrexHistoricData", testBittrexHistoricDataCount)
	t.Run("Data", testDataCount)
	t.Run("PoloniexHistoricData", testPoloniexHistoricDataCount)
	t.Run("Posdatatables", testPosdatatablesCount)
	t.Run("Powdatatables", testPowdatatablesCount)
}

func TestHooks(t *testing.T) {
	t.Run("BittrexHistoricData", testBittrexHistoricDataHooks)
	t.Run("Data", testDataHooks)
	t.Run("PoloniexHistoricData", testPoloniexHistoricDataHooks)
	t.Run("Posdatatables", testPosdatatablesHooks)
	t.Run("Powdatatables", testPowdatatablesHooks)
}

func TestInsert(t *testing.T) {
	t.Run("BittrexHistoricData", testBittrexHistoricDataInsert)
	t.Run("BittrexHistoricData", testBittrexHistoricDataInsertWhitelist)
	t.Run("Data", testDataInsert)
	t.Run("Data", testDataInsertWhitelist)
	t.Run("PoloniexHistoricData", testPoloniexHistoricDataInsert)
	t.Run("PoloniexHistoricData", testPoloniexHistoricDataInsertWhitelist)
	t.Run("Posdatatables", testPosdatatablesInsert)
	t.Run("Posdatatables", testPosdatatablesInsertWhitelist)
	t.Run("Powdatatables", testPowdatatablesInsert)
	t.Run("Powdatatables", testPowdatatablesInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("BittrexHistoricData", testBittrexHistoricDataReload)
	t.Run("Data", testDataReload)
	t.Run("PoloniexHistoricData", testPoloniexHistoricDataReload)
	t.Run("Posdatatables", testPosdatatablesReload)
	t.Run("Powdatatables", testPowdatatablesReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("BittrexHistoricData", testBittrexHistoricDataReloadAll)
	t.Run("Data", testDataReloadAll)
	t.Run("PoloniexHistoricData", testPoloniexHistoricDataReloadAll)
	t.Run("Posdatatables", testPosdatatablesReloadAll)
	t.Run("Powdatatables", testPowdatatablesReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("BittrexHistoricData", testBittrexHistoricDataSelect)
	t.Run("Data", testDataSelect)
	t.Run("PoloniexHistoricData", testPoloniexHistoricDataSelect)
	t.Run("Posdatatables", testPosdatatablesSelect)
	t.Run("Powdatatables", testPowdatatablesSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("BittrexHistoricData", testBittrexHistoricDataUpdate)
	t.Run("Data", testDataUpdate)
	t.Run("PoloniexHistoricData", testPoloniexHistoricDataUpdate)
	t.Run("Posdatatables", testPosdatatablesUpdate)
	t.Run("Powdatatables", testPowdatatablesUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("BittrexHistoricData", testBittrexHistoricDataSliceUpdateAll)
	t.Run("Data", testDataSliceUpdateAll)
	t.Run("PoloniexHistoricData", testPoloniexHistoricDataSliceUpdateAll)
	t.Run("Posdatatables", testPosdatatablesSliceUpdateAll)
	t.Run("Powdatatables", testPowdatatablesSliceUpdateAll)
}

func TestUpsert(t *testing.T) {
	t.Run("BittrexHistoricData", testBittrexHistoricDataUpsert)
	t.Run("Data", testDataUpsert)
	t.Run("PoloniexHistoricData", testPoloniexHistoricDataUpsert)
	t.Run("Posdatatables", testPosdatatablesUpsert)
	t.Run("Powdatatables", testPowdatatablesUpsert)
}
