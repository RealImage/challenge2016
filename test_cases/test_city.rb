require "test/unit"
require_relative "../city"

class Testcity < Test::Unit::TestCase
  def test_all
    assert_equal(0, City.all.count)
    City.create({"Province Code" => "TN", "Country Code" => "IN"})
    assert_equal(0, City.all.count)
    City.create({"City Code" => "CENAI", "City Name" => "Chennai", "Province Code" => "TN", "Country Code" => "IN"})
    assert_equal(1, City.all.count)
  end

  def test_find_by_code
    city = City.create({"City Code" => "TNABE", "City Name" => "Tanabe", "Province Code" => "30", "Country Code" => "JP"})
    assert_equal(City, City.find_by_code("TNABE", "30", "JP").class)
    assert_equal(nil, City.find_by_code("", "30", "JP"))
    assert_equal(nil, City.find_by_code("30"))
    assert_equal(nil, City.find_by_code("TNABE","30","IN"))
    assert_equal(nil, City.find_by_code("TNABE","301","JP"))
  end

  def test_create_state
    city = City.create({"City Code" => "GONCS", "City Name" => "Godmanchester", "Province Code" => "MAN", "Country Code" => "ENG"})
    assert_equal(City, city.class)
    assert_equal("Godmanchester", city.name)
    assert_equal("GONCS", city.code)
    assert_equal("MAN", city.state_code)
    assert_equal("ENG", city.country_code)

    city2 = City.create({"City Code" => "GORM", "City Name" => "Godrome", "Province Code" => "ROM", "Province Name" => "Rome", "Country Code" => "IT"})
    assert_equal(nil, City.create({"City Code" => ""}) )
    assert_equal(nil, City.create({"City Code" => nil}) )
    assert_equal(nil, City.create({"City Code" => "GORM", "City Name" => "Godrome", "Province Code" => "", "Country Code" => "IT"}) )
    assert_equal(nil, City.create({"City Code" => "GORM", "City Name" => "Godrome", "Province Code" => nil, "Country Code" => "IT"}) )
    assert_equal(nil, City.create({"City Code" => "GORM", "City Name" => "Godrome", "Province Code" => "ROM", "Province Name" => "Rome", "Country Code" => ""}) )
    assert_equal(nil, City.create({"City Code" => "GORM", "City Name" => "Godrome", "Province Code" => "ROM", "Province Name" => "Rome", "Country Code" => nil}) )
  end
end