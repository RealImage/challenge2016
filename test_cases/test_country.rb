require "test/unit"
require_relative "../country"

class TestCountry < Test::Unit::TestCase
  def test_all
    Country.create({"Country Code" => "IN", "Country Name" => "India"})
    Country.create({"Country Code" => "US", "Country Name" => "United States"})
    assert_equal(2, Country.all.count)
    Country.create({"Country Code" => nil, "Country Name" => "United States"})
    assert_equal(2, Country.all.count)
    assert_equal([Country],Country.all.to_a.map(&:class).uniq)
  end

  def test_find_by_code
    country = Country.create({"Country Code" => "JP", "Country Name" => "Japan"})
    assert_equal(Country, Country.find_by_code("JP").class)
    assert_equal(nil, Country.find_by_code("PO"))
    assert_equal(nil, Country.find_by_code("PO")&.class)
  end

  def test_create_country
    country = Country.create({"Country Code" => "ENG", "Country Name" => "England"})
    assert_equal(Country, country.class)
    assert_equal("ENG", country.code)
    assert_equal("England", country.name)

    country2 = Country.create({"Country Code" => "IT", "Country Name" => "Italy"})
    assert_equal(nil, Country.create({"Country Code" => "", "Country Name" => "United States"}) )
    assert_equal(nil, Country.create({"Country Code" => nil, "Country Name" => "United States"}) )
    assert_equal(nil, Country.create({"Country Code" => "IT", "Country Name" => "Italy"}) )
  end
end