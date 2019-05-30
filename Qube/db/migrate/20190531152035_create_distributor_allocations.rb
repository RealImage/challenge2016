class CreateDistributorAllocations < ActiveRecord::Migration[5.2]
  def change
    create_table :distributor_allocations do |t|
      t.belongs_to :distributor, foreign_key: true
      t.belongs_to :country, foreign_key: true
      t.belongs_to :province, foreign_key: true
      t.belongs_to :city, foreign_key: true
      t.string :status,default: "included"

      t.timestamps
    end
  end
end
