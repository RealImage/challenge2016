require "test/unit"
require_relative "../distributor"
require_relative "../country"
require_relative "../state"
require_relative "../city"

class TestDistributor < Test::Unit::TestCase
  def test_perform
    assert_equal("Invalid Command", Distributor.perform("abc 123"))
  end

  def test_find_region
    country = Country.create({"Country Code"=>"CN","Country Name"=>"China"})
    assert_equal(country, Distributor.find_region("CN"))
    assert_equal(nil, Distributor.find_region("INI"))
  end

  def test_add
    us = Country.create({"Country Code"=>"US","Country Name"=>"United States"})
    india = Country.create({"Country Code"=>"IN","Country Name"=>"India"})
    tamil_nadu = State.create({"Country Code"=>"IN","Province Code"=>"TN","Province Name"=>"Tamil nadu"})
    karnataka = State.create({"Country Code"=>"IN","Province Code"=>"KA","Province Name"=>"Karnataka"})
    chennai = City.create({"Country Code"=>"IN","Province Code"=>"TN","City Code"=>"CHN", "City Name"=>"Chennai"})
    assert_equal("success",Distributor.add(["Distributor1", "IN,US", "KA-IN,CHN-TN-IN"]))
    assert_equal("Invalid include for distributor",Distributor.add(["Distributor1", "IN,US1", "KA-IN,CHN-TN-IN"]))
    assert_equal("Include should be given",Distributor.add(["Distributor1", "", "KA-IN,CHN-TN-IN"]))
  end

  def test_permission
    us = Country.create({"Country Code"=>"US","Country Name"=>"United States"})
    oklahama = State.create({"Country Code"=>"US","Province Code"=>"OK","Province Name"=>"Oklahama"})
    italy = Country.create({"Country Code"=>"IT","Country Name"=>"Italy"})
    india = Country.create({"Country Code"=>"IN","Country Name"=>"India"})
    tamil_nadu = State.create({"Country Code"=>"IN","Province Code"=>"TN","Province Name"=>"Tamil nadu"})
    karnataka = State.create({"Country Code"=>"IN","Province Code"=>"KA","Province Name"=>"Karnataka"})
    bangalore = City.create({"Country Code"=>"IN","Province Code"=>"KA","City Code"=>"BENAU", "City Name"=>"Bangalore"})
    chennai = City.create({"Country Code"=>"IN","Province Code"=>"TN","City Code"=>"CHN", "City Name"=>"Chennai"})
    chennai = City.create({"Country Code"=>"IN","Province Code"=>"TN","City Code"=>"ED", "City Name"=>"Erode"})
    
    assert_equal("success",Distributor.add(["Distributor1", "IN,US", "KA-IN,CHN-TN-IN"]))
    
    assert_equal("Invalid region", Distributor.permission?(["Distributor1", "BENAU-KA-US"]))
    assert_equal("Distributor not found", Distributor.permission?(["Distributor123", "BENAU-KA-US"]))
    assert_equal("Yes", Distributor.permission?(["Distributor1", "TN-IN"]))
    assert_equal("Yes", Distributor.permission?(["Distributor1", "US"]))
    assert_equal("Yes", Distributor.permission?(["Distributor1", "ED-TN-IN"]))
    assert_equal("No", Distributor.permission?(["Distributor1", "IT"]))
    assert_equal("No", Distributor.permission?(["Distributor1", "CHN-TN-IN"]))
    assert_equal("No", Distributor.permission?(["Distributor1", "KA-IN"]))
    assert_equal("No", Distributor.permission?(["Distributor1", "BENAU-KA-IN"]))
    assert_equal("Yes", Distributor.permission?(["Distributor1", "IN"]))
    assert_equal("Yes", Distributor.permission?(["Distributor1", "OK-US"]))
  end
end