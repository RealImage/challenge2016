class AddAncestryToDistributors < ActiveRecord::Migration
  def change
    add_column :distributors, :ancestry, :string
    add_index :distributors, :ancestry
  end
end
