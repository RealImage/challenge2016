require "test/unit"
require_relative "../state"

class TestState < Test::Unit::TestCase
  def test_all
    State.create({"Province Code" => "TN", "Province Name" => "Tamil Nadu", "Country Code" => "IN"})
    State.create({"Province Code" => "AL", "Province Name" => "Alabama", "Country Code" => "US"})
    assert_equal(2, State.all.count)
    State.create({"Country Code" => nil, "Province Code" => "AL", "Province Name" => "Alabama"})
    State.create({"Country Code" => "", "Province Code" => "AL", "Province Name" => "Alabama"})
    State.create({"Country Code" => "US", "Province Code" => "", "Province Name" => "Alabama"})
    State.create({"Country Code" => "US", "Province Code" => nil, "Province Name" => "Alabama"})
    assert_equal(2, State.all.count)
    assert_equal([State],State.all.to_a.map(&:class).uniq)
  end

  def test_find_by_code
    state = State.create({"Province Code" => "30", "Province Name" => "Wakayama", "Country Code" => "JP"})
    assert_equal(State, State.find_by_code("30", "JP").class)
    assert_equal(nil, State.find_by_code("", "JP"))
    assert_equal(nil, State.find_by_code("30"))
    assert_equal(nil, State.find_by_code("30","IN"))
  end

  def test_create_state
    state = State.create({"Province Code" => "MAN", "Province Name" => "Manchester", "Country Code" => "ENG"})
    assert_equal(State, state.class)
    assert_equal("MAN", state.code)
    assert_equal("Manchester", state.name)
    assert_equal("ENG", state.country_code)

    state2 = State.create({"Province Code" => "ROM", "Province Name" => "Rome","Country Code" => "IT"})
    assert_equal(nil, State.create({"Country Code" => ""}) )
    assert_equal(nil, State.create({"Country Code" => nil}) )
    assert_equal(nil, State.create({"Province Code" => "", "Country Code" => "IT"}) )
    assert_equal(nil, State.create({"Province Code" => nil, "Country Code" => "IT"}) )
  end
end