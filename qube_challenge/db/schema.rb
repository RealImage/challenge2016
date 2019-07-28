# encoding: UTF-8
# This file is auto-generated from the current state of the database. Instead
# of editing this file, please use the migrations feature of Active Record to
# incrementally modify your database, and then regenerate this schema definition.
#
# Note that this schema.rb definition is the authoritative source for your
# database schema. If you need to create the application database on another
# system, you should be using db:schema:load, not running all the migrations
# from scratch. The latter is a flawed and unsustainable approach (the more migrations
# you'll amass, the slower it'll run and the greater likelihood for issues).
#
# It's strongly recommended that you check this file into your version control system.

ActiveRecord::Schema.define(version: 20190629071102) do

  # These are extensions that must be enabled in order to support this database
  enable_extension "plpgsql"

  create_table "cities", force: :cascade do |t|
    t.string   "name"
    t.string   "code"
    t.string   "province_code"
    t.datetime "created_at",    null: false
    t.datetime "updated_at",    null: false
  end

  create_table "countries", force: :cascade do |t|
    t.string   "name"
    t.string   "code"
    t.datetime "created_at", null: false
    t.datetime "updated_at", null: false
  end

  create_table "distributor_areas", force: :cascade do |t|
    t.integer  "distributor_id"
    t.string   "country_code"
    t.string   "province_code"
    t.string   "city_code"
    t.boolean  "is_included"
    t.datetime "created_at",     null: false
    t.datetime "updated_at",     null: false
  end

  add_index "distributor_areas", ["distributor_id"], name: "index_distributor_areas_on_distributor_id", using: :btree

  create_table "distributors", force: :cascade do |t|
    t.string   "name"
    t.integer  "parent_id"
    t.datetime "created_at", null: false
    t.datetime "updated_at", null: false
    t.string   "ancestry"
  end

  add_index "distributors", ["ancestry"], name: "index_distributors_on_ancestry", using: :btree

  create_table "provinces", force: :cascade do |t|
    t.string   "name"
    t.string   "code"
    t.string   "country_code"
    t.datetime "created_at",   null: false
    t.datetime "updated_at",   null: false
  end

  add_foreign_key "distributor_areas", "distributors"
end
