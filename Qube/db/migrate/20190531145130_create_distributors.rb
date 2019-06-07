class CreateDistributors < ActiveRecord::Migration[5.2]
  def change
  	create_table :distributors do |t|
      t.string  :name
      t.integer :parent_id
      t.string  :status, default: "active"
      t.timestamp
    end
    
    add_index :distributors,:parent_id,:name => "index_parent_id"
  end
end