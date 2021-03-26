import { ApolloServer } from 'apollo-server/dist'
import { ApolloGateway } from '@apollo/gateway'

type Service = {
    name: string
    url: string
}

const serviceList: Service[] = [
  { name: 'accounts', url: process.env.ACCOUNTS_URL !== undefined ? process.env.ACCOUNTS_URL : 'http://localhost:4001/query' },
  { name: 'products', url: process.env.PRODUCTS_URL !== undefined ? process.env.PRODUCTS_URL : 'http://localhost:4002/query' },
  { name: 'reviews', url: process.env.REVIEWS_URL !== undefined ? process.env.REVIEWS_URL : 'http://localhost:4003/query' }
]

const gateway = new ApolloGateway({
  serviceList
})

const server = new ApolloServer({
  gateway,

  subscriptions: false
})

server.listen().then(({ url }: {url: string}) => {
  console.log(`🚀 Server ready at ${url}`)
})
