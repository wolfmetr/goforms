package forms_test

import (
	. "launchpad.net/gocheck"
	"launchpad.net/goforms"
)

type CharFieldTestSuite struct{}

var _ = Suite(&CharFieldTestSuite{})

func (s *CharFieldTestSuite) TestCleanSuccess(c *C) {
	f := forms.NewCharField("description")
	f.SetValue("Testing 1, 2, 3")

	c.Check(f.Clean(), Equals, nil)
	c.Check(f.CleanedValue(), Equals, "Testing 1, 2, 3")
}

func (s *CharFieldTestSuite) TestMinLength(c *C) {
	f := forms.NewCharField("description")
	f.MinLength = 20
	f.SetValue("Testing 1, 2, 3")

	err := f.Clean()

	c.Check(err.String(), Equals,
		"The value must have a minimum length of 20 characters.")
}

func (s *CharFieldTestSuite) TestMaxLength(c *C) {
	f := forms.NewCharField("description")
	f.MaxLength = 10
	f.SetValue("Testing 1, 2, 3")

	err := f.Clean()

	c.Check(err.String(), Equals,
		"The value must have a maximum length of 10 characters.")
}

type IntegerFieldTestSuite struct{}

var _ = Suite(&IntegerFieldTestSuite{})

func (s *IntegerFieldTestSuite) TestCleanSuccess(c *C) {
	f := forms.NewIntegerField("num_purchases")
	f.SetValue("12345")

	c.Check(f.Clean(), Equals, nil)
	c.Check(f.CleanedValue(), Equals, 12345)
}

func (s *IntegerFieldTestSuite) TestCleanInvalid(c *C) {
	f := forms.NewIntegerField("num_purchases")
	f.SetValue("a12345")

	err := f.Clean()
	c.Check(err.String(), Equals, "The value must be a valid integer.")
}

type RegexFieldTestSuite struct{}

var _ = Suite(&RegexFieldTestSuite{})

func (s *RegexFieldTestSuite) TestCleanSuccess(c *C) {
	f := forms.NewRegexField("alphabet", "a.c")
	f.SetValue("abc")

	c.Check(f.Clean(), Equals, nil)
	c.Check(f.CleanedValue(), Equals, "abc")
}

func (s *RegexFieldTestSuite) TestCleanInvalid(c *C) {
	f := forms.NewRegexField("alphabet", "a.c")
	f.SetValue("abz")

	err := f.Clean()
	c.Check(err.String(), Equals, "The input 'abz' did not match 'a.c'.")
	c.Check(f.CleanedValue(), IsNil)
}