package assert

import (
	"time"

	"github.com/smartystreets/assertions"
	"github.com/smartystreets/assertions/should"
)

// Tha reads like 'That' (so long as you are passing a *testing.T named 't'.
//
// For example:
//
//	assert.Tha(t, 42).IsBetween(0, 100)
//
// ...which correlates very closely with the following prose:
//
//	"Assert that 42 is between 0 and 100."
func Tha(t testingT, actual interface{}) *assertion {
	t.Helper()
	return &assertion{
		t:      t,
		actual: actual,
	}
}

type testingT interface {
	Helper()
	Error(args ...interface{})
}

type assertion struct {
	t      testingT
	actual interface{}
}

func (this *assertion) so(actual interface{}, assert assertions.SoFunc, expected ...interface{}) {
	this.t.Helper()
	ok, result := assertions.So(actual, assert, expected...)
	if !ok {
		this.t.Error("\n" + result)
	}
}

// AlmostEquals is analogous to should.AlmostEqual.
func (this *assertion) AlmostEquals(expected interface{}, tolerance float64) {
	this.t.Helper()
	this.so(this.actual, should.AlmostEqual, expected, tolerance)
}

// IsBetween is analogous to should.BeBetween.
func (this *assertion) IsBetween(lower, upper interface{}) {
	this.t.Helper()
	this.so(this.actual, should.BeBetween, lower, upper)
}

// IsBetweenOrEqual is analogous to should.BeBetweenOrEqual.
func (this *assertion) IsBetweenOrEqual(lower, upper interface{}) {
	this.t.Helper()
	this.so(this.actual, should.BeBetweenOrEqual, lower, upper)
}

// IsBlank is analogous to should.BeBlank.
func (this *assertion) IsBlank() {
	this.t.Helper()
	this.so(this.actual, should.BeBlank)
}

// AreChronological is analogous to should.BeChronological.
func (this *assertion) AreChronological() {
	this.t.Helper()
	this.so(this.actual, should.BeChronological)
}

// IsEmpty is analogous to should.BeEmpty.
func (this *assertion) IsEmpty() {
	this.t.Helper()
	this.so(this.actual, should.BeEmpty)
}

// IsError is analogous to should.BeError.
func (this *assertion) IsError(expected ...interface{}) {
	this.t.Helper()
	this.so(this.actual, should.BeError, expected...)
}

// IsFalse is analogous to should.BeFalse.
func (this *assertion) IsFalse() {
	this.t.Helper()
	this.so(this.actual, should.BeFalse)
}

// IsGreaterThan is analogous to should.BeGreaterThan.
func (this *assertion) IsGreaterThan(expected interface{}) {
	this.t.Helper()
	this.so(this.actual, should.BeGreaterThan, expected)
}

// IsGreaterThanOrEqualTo is analogous to should.BeGreaterThanOrEqualTo.
func (this *assertion) IsGreaterThanOrEqualTo(expected interface{}) {
	this.t.Helper()
	this.so(this.actual, should.BeGreaterThanOrEqualTo, expected)
}

// IsIn is analogous to should.BeIn.
func (this *assertion) IsIn(expected interface{}) {
	this.t.Helper()
	this.so(this.actual, should.BeIn, expected)
}

// IsLessThan is analogous to should.BeLessThan.
func (this *assertion) IsLessThan(expected interface{}) {
	this.t.Helper()
	this.so(this.actual, should.BeLessThan, expected)
}

// IsLessThanOrEqualTo is analogous to should.BeLessThanOrEqualTo.
func (this *assertion) IsLessThanOrEqualTo(expected interface{}) {
	this.t.Helper()
	this.so(this.actual, should.BeLessThanOrEqualTo, expected)
}

// IsNil is analogous to should.BeNil.
func (this *assertion) IsNil() {
	this.t.Helper()
	this.so(this.actual, should.BeNil)
}

// IsTrue is analogous to should.BeTrue.
func (this *assertion) IsTrue() {
	this.t.Helper()
	this.so(this.actual, should.BeTrue)
}

// IsZeroValue is analogous to should.BeZeroValue.
func (this *assertion) IsZeroValue() {
	this.t.Helper()
	this.so(this.actual, should.BeZeroValue)
}

// Contains is analogous to should.Contain.
func (this *assertion) Contains(expected interface{}) {
	this.t.Helper()
	this.so(this.actual, should.Contain, expected)
}

// ContainsKey is analogous to should.ContainKey.
func (this *assertion) ContainsKey(expected interface{}) {
	this.t.Helper()
	this.so(this.actual, should.ContainKey, expected)
}

// ContainsSubstring is analogous to should.ContainSubstring.
func (this *assertion) ContainsSubstring(expected string) {
	this.t.Helper()
	this.so(this.actual, should.ContainSubstring, expected)
}

// EndsWith is analogous to should.EndWith.
func (this *assertion) EndsWith(expected string) {
	this.t.Helper()
	this.so(this.actual, should.EndWith, expected)
}

// Equals is analogous to should.Equal.
func (this *assertion) Equals(expected interface{}) {
	this.t.Helper()
	this.so(this.actual, should.Equal, expected)
}

// EqualsJSON is analogous to should.EqualJSON.
func (this *assertion) EqualsJSON(expected string) {
	this.t.Helper()
	this.so(this.actual, should.EqualJSON, expected)
}

// EqualsTrimSpace is analogous to should.EqualTrimSpace.
func (this *assertion) EqualsTrimSpace(expected string) {
	this.t.Helper()
	this.so(this.actual, should.EqualTrimSpace, expected)
}

// EqualsWithout is analogous to should.EqualWithout.
func (this *assertion) EqualsWithout(expected, remove string) {
	this.t.Helper()
	this.so(this.actual, should.EqualWithout, expected, remove)
}

// HappensAfter is analogous to should.HappenAfter.
func (this *assertion) HappensAfter(expected time.Time) {
	this.t.Helper()
	this.so(this.actual, should.HappenAfter, expected)
}

// HappensBefore is analogous to should.HappenBefore.
func (this *assertion) HappensBefore(expected time.Time) {
	this.t.Helper()
	this.so(this.actual, should.HappenBefore, expected)
}

// HappensBetween is analogous to should.HappenBetween.
func (this *assertion) HappensBetween(before, after time.Time) {
	this.t.Helper()
	this.so(this.actual, should.HappenBetween, before, after)
}

// HappensOnOrAfter is analogous to should.HappenOnOrAfter.
func (this *assertion) HappensOnOrAfter(expected time.Time) {
	this.t.Helper()
	this.so(this.actual, should.HappenOnOrAfter, expected)
}

// HappensOnOrBefore is analogous to should.HappenOnOrBefore.
func (this *assertion) HappensOnOrBefore(expected time.Time) {
	this.t.Helper()
	this.so(this.actual, should.HappenOnOrBefore, expected)
}

// HappensOnOrBetween is analogous to should.HappenOnOrBetween.
func (this *assertion) HappensOnOrBetween(before, after time.Time) {
	this.t.Helper()
	this.so(this.actual, should.HappenOnOrBetween, before, after)
}

// HappensWithin is analogous to should.HappenWithin.
func (this *assertion) HappensWithin(duration time.Duration, after time.Time) {
	this.t.Helper()
	this.so(this.actual, should.HappenWithin, duration, after)
}

// HasLength is analogous to should.HaveLength.
func (this *assertion) HasLength(expected interface{}) {
	this.t.Helper()
	this.so(this.actual, should.HaveLength, expected)
}

// HasSameTypeAs is analogous to should.HaveSameTypeAs.
func (this *assertion) HasSameTypeAs(expected interface{}) {
	this.t.Helper()
	this.so(this.actual, should.HaveSameTypeAs, expected)
}

// Implements is analogous to should.Implement.
func (this *assertion) Implements(expected interface{}) {
	this.t.Helper()
	this.so(this.actual, should.Implement, expected)
}

// IsNotAlmostEqual is analogous to should.NotAlmostEqual.
func (this *assertion) IsNotAlmostEqual(expected interface{}, tolerance float64) {
	this.t.Helper()
	this.so(this.actual, should.NotAlmostEqual, expected, tolerance)
}

// IsNotBetween is analogous to should.NotBeBetween.
func (this *assertion) IsNotBetween(lower, upper interface{}) {
	this.t.Helper()
	this.so(this.actual, should.NotBeBetween, lower, upper)
}

// IsNotBetweenOrEqual is analogous to should.NotBeBetweenOrEqual.
func (this *assertion) IsNotBetweenOrEqual(lower, upper interface{}) {
	this.t.Helper()
	this.so(this.actual, should.NotBeBetweenOrEqual, lower, upper)
}

// IsNotBlank is analogous to should.NotBeBlank.
func (this *assertion) IsNotBlank() {
	this.t.Helper()
	this.so(this.actual, should.NotBeBlank)
}

// AreNotChronological is analogous to should.NotBeChronological.
func (this *assertion) AreNotChronological() {
	this.t.Helper()
	this.so(this.actual, should.NotBeChronological)
}

// IsNotEmpty is analogous to should.NotBeEmpty.
func (this *assertion) IsNotEmpty() {
	this.t.Helper()
	this.so(this.actual, should.NotBeEmpty)
}

// IsNotIn is analogous to should.NotBeIn.
func (this *assertion) IsNotIn(expected interface{}) {
	this.t.Helper()
	this.so(this.actual, should.NotBeIn, expected)
}

// IsNotNil is analogous to should.NotBeNil.
func (this *assertion) IsNotNil() {
	this.t.Helper()
	this.so(this.actual, should.NotBeNil)
}

// IsNotZeroValue is analogous to should.NotBeZeroValue.
func (this *assertion) IsNotZeroValue() {
	this.t.Helper()
	this.so(this.actual, should.NotBeZeroValue)
}

// DoesNotContain is analogous to should.NotContain.
func (this *assertion) DoesNotContain(expected interface{}) {
	this.t.Helper()
	this.so(this.actual, should.NotContain, expected)
}

// DoesNotContainKey is analogous to should.NotContainKey.
func (this *assertion) DoesNotContainKey(expected interface{}) {
	this.t.Helper()
	this.so(this.actual, should.NotContainKey, expected)
}

// DoesNotContainSubstring is analogous to should.NotContainSubstring.
func (this *assertion) DoesNotContainSubstring(expected interface{}) {
	this.t.Helper()
	this.so(this.actual, should.NotContainSubstring, expected)
}

// DoesNotEndWith is analogous to should.NotEndWith.
func (this *assertion) DoesNotEndWith(expected interface{}) {
	this.t.Helper()
	this.so(this.actual, should.NotEndWith, expected)
}

// DoesNotEqual is analogous to should.NotEqual.
func (this *assertion) DoesNotEqual(expected interface{}) {
	this.t.Helper()
	this.so(this.actual, should.NotEqual, expected)
}

// DoesNotHappenOnOrBetween is analogous to should.NotHappenOnOrBetween.
func (this *assertion) DoesNotHappenOnOrBetween(before, after time.Time) {
	this.t.Helper()
	this.so(this.actual, should.NotHappenOnOrBetween, before, after)
}

// DoesNotHappenWithin is analogous to should.NotHappenWithin.
func (this *assertion) DoesNotHappenWithin(duration time.Duration, expected time.Time) {
	this.t.Helper()
	this.so(this.actual, should.NotHappenWithin, duration, expected)
}

// DoesNotHaveSameTypeAs is analogous to should.NotHaveSameTypeAs.
func (this *assertion) DoesNotHaveSameTypeAs(expected interface{}) {
	this.t.Helper()
	this.so(this.actual, should.NotHaveSameTypeAs, expected)
}

// DoesNotImplement is analogous to should.NotImplement.
func (this *assertion) DoesNotImplement(expected interface{}) {
	this.t.Helper()
	this.so(this.actual, should.NotImplement, expected)
}

// DoesNotPanic is analogous to should.NotPanic.
func (this *assertion) DoesNotPanic() {
	this.t.Helper()
	this.so(this.actual, should.NotPanic)
}

// DoesNotPanicWith is analogous to should.NotPanicWith.
func (this *assertion) DoesNotPanicWith(expected interface{}) {
	this.t.Helper()
	this.so(this.actual, should.NotPanicWith, expected)
}

// DoesNotPointTo is analogous to should.NotPointTo.
func (this *assertion) DoesNotPointTo(expected interface{}) {
	this.t.Helper()
	this.so(this.actual, should.NotPointTo, expected)
}

// DoesNotResemble is analogous to should.NotResemble.
func (this *assertion) DoesNotResemble(expected interface{}) {
	this.t.Helper()
	this.so(this.actual, should.NotResemble, expected)
}

// DoesNotStartWith is analogous to should.NotStartWith.
func (this *assertion) DoesNotStartWith(expected interface{}) {
	this.t.Helper()
	this.so(this.actual, should.NotStartWith, expected)
}

// Panics is analogous to should.Panic.
func (this *assertion) Panics() {
	this.t.Helper()
	this.so(this.actual, should.Panic)
}

// PanicsWith is analogous to should.PanicWith.
func (this *assertion) PanicsWith(expected interface{}) {
	this.t.Helper()
	this.so(this.actual, should.PanicWith, expected)
}

// PointsTo is analogous to should.PointTo.
func (this *assertion) PointsTo(expected interface{}) {
	this.t.Helper()
	this.so(this.actual, should.PointTo, expected)
}

// Resembles is analogous to should.Resemble.
func (this *assertion) Resembles(expected interface{}) {
	this.t.Helper()
	this.so(this.actual, should.Resemble, expected)
}

// StartsWith is analogous to should.StartWith.
func (this *assertion) StartsWith(expected interface{}) {
	this.t.Helper()
	this.so(this.actual, should.StartWith, expected)
}

// Wraps is analogous to should.Wrap.
func (this *assertion) Wraps(expected interface{}) {
	this.t.Helper()
	this.so(this.actual, should.Wrap, expected)
}
