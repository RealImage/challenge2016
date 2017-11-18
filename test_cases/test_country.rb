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
    country = Country.create({"Country Code" => "IN", "Country Name" => "India"})
    assert_equal("Country", country.class.name)
    assert_equal("IN", country.code)
    assert_equal("India", country.name)

    country2 = Country.create({"Country Code" => "US", "Country Name" => "United States"})
    assert_equal(nil, Country.create({"Country Code" => "", "Country Name" => "United States"}) )
    assert_equal(nil, Country.create({"Country Code" => nil, "Country Name" => "United States"}) )
  end
end