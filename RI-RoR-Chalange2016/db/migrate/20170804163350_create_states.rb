class CreateStates < ActiveRecord::Migration[5.0]
  def change
    create_table :states do |t|
      t.string :name
      t.string :code
      t.references :country

      t.timestamps
    end
  end
end
