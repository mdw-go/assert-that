package assert_test

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/mdwhatcott/assert-that/assert"
)

func TestAlmostEquals(t *testing.T)             { assert.Tha(t, 42).AlmostEquals(42.001, .01) }
func TestIsBetween(t *testing.T)                { assert.Tha(t, 42).IsBetween(0, 100) }
func TestIsBetweenOrEqual(t *testing.T)         { assert.Tha(t, 42).IsBetweenOrEqual(42, 43) }
func TestIsBlank(t *testing.T)                  { assert.Tha(t, "").IsBlank() }
func TestAreChronological(t *testing.T)         { assert.Tha(t, orderedTimes).AreChronological() }
func TestIsEmpty(t *testing.T)                  { assert.Tha(t, []int{}).IsEmpty() }
func TestIsError(t *testing.T)                  { assert.Tha(t, errHi).IsError() }
func TestIsFalse(t *testing.T)                  { assert.Tha(t, false).IsFalse() }
func TestIsGreaterThan(t *testing.T)            { assert.Tha(t, 42).IsGreaterThan(41) }
func TestIsGreaterThanOrEqualTo(t *testing.T)   { assert.Tha(t, 42).IsGreaterThanOrEqualTo(42) }
func TestIsIn(t *testing.T)                     { assert.Tha(t, 42).IsIn([]int{42}) }
func TestIsLessThan(t *testing.T)               { assert.Tha(t, 42).IsLessThan(43) }
func TestIsLessThanOrEqualTo(t *testing.T)      { assert.Tha(t, 42).IsLessThanOrEqualTo(42) }
func TestIsNil(t *testing.T)                    { assert.Tha(t, nil).IsNil() }
func TestIsTrue(t *testing.T)                   { assert.Tha(t, true).IsTrue() }
func TestIsZeroValue(t *testing.T)              { assert.Tha(t, 0).IsZeroValue() }
func TestContains(t *testing.T)                 { assert.Tha(t, []int{42}).Contains(42) }
func TestContainsKey(t *testing.T)              { assert.Tha(t, map[int]int{42: 1}).ContainsKey(42) }
func TestContainsSubstring(t *testing.T)        { assert.Tha(t, "hello").ContainsSubstring("ll") }
func TestEndsWith(t *testing.T)                 { assert.Tha(t, "hello").EndsWith("lo") }
func TestEquals(t *testing.T)                   { assert.Tha(t, 42).Equals(42) }
func TestEqualsJSON(t *testing.T)               { assert.Tha(t, "42").EqualsJSON("42") }
func TestEqualsTrimSpace(t *testing.T)          { assert.Tha(t, " 42 ").EqualsTrimSpace("42") }
func TestEqualsWithout(t *testing.T)            { assert.Tha(t, "42").EqualsWithout("4", "2") }
func TestHappensAfter(t *testing.T)             { assert.Tha(t, t3).HappensAfter(t1) }
func TestHappensBefore(t *testing.T)            { assert.Tha(t, t1).HappensBefore(t3) }
func TestHappensBetween(t *testing.T)           { assert.Tha(t, t2).HappensBetween(t1, t3) }
func TestHappensOnOrAfter(t *testing.T)         { assert.Tha(t, t2).HappensOnOrAfter(t1) }
func TestHappensOnOrBefore(t *testing.T)        { assert.Tha(t, t2).HappensOnOrBefore(t3) }
func TestHappensOnOrBetween(t *testing.T)       { assert.Tha(t, t2).HappensOnOrBetween(t1, t3) }
func TestHappensWithin(t *testing.T)            { assert.Tha(t, t2).HappensWithin(time.Second, t3) }
func TestHasLength(t *testing.T)                { assert.Tha(t, []int{1, 2, 3}).HasLength(3) }
func TestHasSameTypeAs(t *testing.T)            { assert.Tha(t, 42).HasSameTypeAs(42) }
func TestImplements(t *testing.T)               { assert.Tha(t, errHi).Implements((*error)(nil)) }
func TestIsNotAlmostEqual(t *testing.T)         { assert.Tha(t, 42).IsNotAlmostEqual(43, .01) }
func TestIsNotBetween(t *testing.T)             { assert.Tha(t, 42).IsNotBetween(43, 44) }
func TestIsNotBetweenOrEqual(t *testing.T)      { assert.Tha(t, 42).IsNotBetweenOrEqual(43, 44) }
func TestIsNotBlank(t *testing.T)               { assert.Tha(t, "hi").IsNotBlank() }
func TestIsNotChronological(t *testing.T)       { assert.Tha(t, unorderedTimes).AreNotChronological() }
func TestIsNotEmpty(t *testing.T)               { assert.Tha(t, []int{1}).IsNotEmpty() }
func TestIsNotIn(t *testing.T)                  { assert.Tha(t, 42).IsNotIn(42) }
func TestIsNotNil(t *testing.T)                 { assert.Tha(t, errHi).IsNotNil() }
func TestIsNotZeroValue(t *testing.T)           { assert.Tha(t, 42).IsNotZeroValue() }
func TestDoesNotContain(t *testing.T)           { assert.Tha(t, []int{1}).DoesNotContain(42) }
func TestDoesNotContainKey(t *testing.T)        { assert.Tha(t, map[int]int{1: 2}).DoesNotContainKey(42) }
func TestDoesNotContainSubstring(t *testing.T)  { assert.Tha(t, "42").DoesNotContainSubstring("3") }
func TestDoesNotEndWith(t *testing.T)           { assert.Tha(t, "42").DoesNotEndWith("3") }
func TestDoesNotEqual(t *testing.T)             { assert.Tha(t, 42).DoesNotEqual(43) }
func TestDoesNotHappenOnOrBetween(t *testing.T) { assert.Tha(t, t3).DoesNotHappenOnOrBetween(t1, t2) }
func TestDoesNotHappenWithin(t *testing.T)      { assert.Tha(t, t2).DoesNotHappenWithin(1, t3) }
func TestDoesNotHaveSameTypeAs(t *testing.T)    { assert.Tha(t, 42).DoesNotHaveSameTypeAs(uint(42)) }
func TestDoesNotImplement(t *testing.T)         { assert.Tha(t, 42).DoesNotImplement((*error)(nil)) }
func TestDoesNotPanic(t *testing.T)             { assert.Tha(t, func() {}).DoesNotPanic() }
func TestDoesNotPanicWith(t *testing.T)         { assert.Tha(t, func() {}).DoesNotPanicWith(42) }
func TestDoesNotPointTo(t *testing.T)           { assert.Tha(t, &time.Time{}).DoesNotPointTo(&time.Time{}) }
func TestDoesNotResemble(t *testing.T)          { assert.Tha(t, 42).DoesNotResemble(43) }
func TestDoesNotStartWith(t *testing.T)         { assert.Tha(t, "42").DoesNotStartWith("3") }
func TestPanics(t *testing.T)                   { assert.Tha(t, func() { panic(42) }).Panics() }
func TestPanicsWith(t *testing.T)               { assert.Tha(t, func() { panic(42) }).PanicsWith(42) }
func TestPointsTo(t *testing.T)                 { assert.Tha(t, &t2).PointsTo(&t2) }
func TestResembles(t *testing.T)                { assert.Tha(t, 42).Resembles(42) }
func TestStartsWith(t *testing.T)               { assert.Tha(t, "42").StartsWith("42") }
func TestWraps(t *testing.T)                    { assert.Tha(t, fmt.Errorf("%w", errHi)).Wraps(errHi) }

var (
	orderedTimes   = []time.Time{time.Now(), time.Now(), time.Now()}
	unorderedTimes = []time.Time{t3, t1, t2}

	t1 = t2.Add(-time.Second)
	t2 = time.Now()
	t3 = t2.Add(time.Second)

	errHi = errors.New("hi")
)

type FakeTestingT struct{ failed bool }

func (this *FakeTestingT) Helper()              {}
func (this *FakeTestingT) Error(...interface{}) { this.failed = true }

func TestFailure(t *testing.T) {
	T := new(FakeTestingT)
	assert.Tha(T, false).IsTrue()
	assert.Tha(t, T.failed).IsTrue()
}
