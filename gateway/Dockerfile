FROM node:14-slim as build

WORKDIR /app

COPY package.json yarn.lock ./

RUN yarn install --non-interactive --frozen-lockfile

COPY . .

RUN yarn build

FROM node:14-slim as node_modules

WORKDIR /app

COPY package.json yarn.lock ./

RUN yarn install --non-interactive --frozen-lockfile --prod

FROM gcr.io/distroless/nodejs:14
COPY . /app
WORKDIR /app

ENV NODE_ENV=production

USER nonroot
COPY --chown=nonroot:nonroot --from=build /app/dist /app/dist
COPY --chown=nonroot:nonroot --from=node_modules /app/node_modules /app/node_modules

CMD ["dist/index.js"]
