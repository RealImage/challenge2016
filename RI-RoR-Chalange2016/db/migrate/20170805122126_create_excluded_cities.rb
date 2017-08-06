class CreateExcludedCities < ActiveRecord::Migration[5.0]
  def change
    create_table :excluded_cities do |t|
      t.references :distbutor, foreign_key: true
      t.references :city, foreign_key: true

      t.timestamps
    end
  end
end
