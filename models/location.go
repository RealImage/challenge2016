package models

type Cities []string

type Province map[string]Cities // map of cities aka province

type Location map[string][]Province // map of provinces aka country
