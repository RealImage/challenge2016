class CreateIncludedCountries < ActiveRecord::Migration[5.0]
  def change
    create_table :included_countries do |t|
      t.references :distbutor, foreign_key: true
      t.references :country, foreign_key: true

      t.timestamps
    end
  end
end
