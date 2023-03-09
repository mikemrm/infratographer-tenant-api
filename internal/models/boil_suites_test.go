// Code generated by SQLBoiler 4.14.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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
	t.Run("Tenants", testTenants)
}

func TestSoftDelete(t *testing.T) {
	t.Run("Tenants", testTenantsSoftDelete)
}

func TestQuerySoftDeleteAll(t *testing.T) {
	t.Run("Tenants", testTenantsQuerySoftDeleteAll)
}

func TestSliceSoftDeleteAll(t *testing.T) {
	t.Run("Tenants", testTenantsSliceSoftDeleteAll)
}

func TestDelete(t *testing.T) {
	t.Run("Tenants", testTenantsDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Tenants", testTenantsQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Tenants", testTenantsSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Tenants", testTenantsExists)
}

func TestFind(t *testing.T) {
	t.Run("Tenants", testTenantsFind)
}

func TestBind(t *testing.T) {
	t.Run("Tenants", testTenantsBind)
}

func TestOne(t *testing.T) {
	t.Run("Tenants", testTenantsOne)
}

func TestAll(t *testing.T) {
	t.Run("Tenants", testTenantsAll)
}

func TestCount(t *testing.T) {
	t.Run("Tenants", testTenantsCount)
}

func TestHooks(t *testing.T) {
	t.Run("Tenants", testTenantsHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Tenants", testTenantsInsert)
	t.Run("Tenants", testTenantsInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("TenantToTenantUsingParentTenant", testTenantToOneTenantUsingParentTenant)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("TenantToParentTenantTenants", testTenantToManyParentTenantTenants)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("TenantToTenantUsingParentTenantTenants", testTenantToOneSetOpTenantUsingParentTenant)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {
	t.Run("TenantToTenantUsingParentTenantTenants", testTenantToOneRemoveOpTenantUsingParentTenant)
}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("TenantToParentTenantTenants", testTenantToManyAddOpParentTenantTenants)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {
	t.Run("TenantToParentTenantTenants", testTenantToManySetOpParentTenantTenants)
}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {
	t.Run("TenantToParentTenantTenants", testTenantToManyRemoveOpParentTenantTenants)
}

func TestReload(t *testing.T) {
	t.Run("Tenants", testTenantsReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Tenants", testTenantsReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Tenants", testTenantsSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Tenants", testTenantsUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Tenants", testTenantsSliceUpdateAll)
}