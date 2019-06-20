module Simulations
  module Config
    def distributors
      [
        {
          name: 'DISTRIBUTOR1',
          included: ['IN'],
          excluded: ['PUNCH,JK,IN'],
          sub_distributors: [
            {
              name: 'DISTRIBUTOR2',
              included: ['TN,IN'],
              excluded: ['KA,IN']
            }
          ]
        },
        {
          name: 'DISTRIBUTOR3',
          included: ['IN'],
          excluded: ['KA,IN'],
          sub_distributors: [
            {
              name: 'DISTRIBUTOR4',
              included: ['TN,IN', 'KA,IN', 'SOWAE,KA,IN'],
              excluded: ['KNGLM,TN,IN'],
              sub_distributors: [
                {
                  name: 'DISTRIBUTOR5',
                  included: [],
                  excluded: []
                }
              ]
            }
          ]
        },
        {
          name: 'DISTRIBUTOR7',
          included: [],
          excluded: [],
          sub_distributors: [
            {
              name: 'DISTRIBUTOR8',
              included: [],
              excluded: []
            }
          ]
        }
      ]
    end
  end
end
