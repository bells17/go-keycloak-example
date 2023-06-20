class CreateAccounts < ActiveRecord::Migration[6.1]
  def change
    create_table :accounts do |t|

      t.timestamps
    end
  end
end
