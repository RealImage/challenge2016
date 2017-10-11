class CreateExcludedStates < ActiveRecord::Migration[5.0]
  def change
    create_table :excluded_states do |t|
      t.references :distbutor, foreign_key: true
      t.references :state, foreign_key: true

      t.timestamps
    end
  end
end
