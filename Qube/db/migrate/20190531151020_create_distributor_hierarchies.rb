class CreateDistributorHierarchies < ActiveRecord::Migration[5.2]
  def change
    create_table :distributor_hierarchies, id: false do |t|
      t.integer :ancestor_id, null: false
      t.integer :descendant_id, null: false
      t.integer :generations, null: false
    end

    add_index :distributor_hierarchies, [:ancestor_id, :descendant_id, :generations],
      unique: true,
      name: "distributor_anc_desc_idx"

    add_index :distributor_hierarchies, [:descendant_id],
      name: "distributor_desc_idx"
  end
end
