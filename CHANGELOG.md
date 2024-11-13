# Changelog

All notable changes to this project will be documented in this file. See [conventional commits](https://www.conventionalcommits.org/) for commit guidelines.

---
## [unreleased]

### Bug Fixes

- **(api)** use db package from api and not consumer - ([8121a0b](https://github.com/cocogitto/cocogitto/commit/8121a0be518ea715113b20f54dcf5e154bad756e)) - christoff-linde
- **(codeql)** remove dummy echo commands - ([400aab3](https://github.com/cocogitto/cocogitto/commit/400aab3e12d6ba328e6b5ca84ef68cbcdb1193a8)) - christoff-linde
- **(consumer)** block writes of 0 values in sensor_reading - ([bd03587](https://github.com/cocogitto/cocogitto/commit/bd03587cf30d61d1b1da2d21dd0913890654514b)) - christoff-linde
- **(db)** ensure sensor_id is actual FK in sensor_metadata - ([74b12e6](https://github.com/cocogitto/cocogitto/commit/74b12e6ea7a7f080a09ada2cbf1450e5a81137e0)) - christoff-linde
- **(justfile)** run command for consumer app - ([9d3cfb8](https://github.com/cocogitto/cocogitto/commit/9d3cfb85118ca81b73399394480abaebc7ca6037)) - christoff-linde

### Documentation

- **(.env.example)** add sample env file - ([9c61322](https://github.com/cocogitto/cocogitto/commit/9c613221b651ca56c5e649e5f873b1b2d159ae4e)) - christoff-linde
- **(.env.example)** add db_url to example - ([5f4808d](https://github.com/cocogitto/cocogitto/commit/5f4808dd75e7047d25c44d283bcbf9c9ad4e94a9)) - christoff-linde
- **(README)** rename project & update description - ([c4a85f5](https://github.com/cocogitto/cocogitto/commit/c4a85f52a90a170744289612a5ae6d0ba4b0d5e8)) - christoff-linde
- **(README)** add mermaid doc for db schema - ([5be26ec](https://github.com/cocogitto/cocogitto/commit/5be26ec8def26b654e52cb84075d49929fada8f7)) - Christoff Linde
- **(db)** add project structure description document - ([54742bb](https://github.com/cocogitto/cocogitto/commit/54742bbf9917caeace22405abaf5e2ea51f2fec7)) - christoff-linde

### Features

- **(.env.example)** add API_PORT to example - ([118a70c](https://github.com/cocogitto/cocogitto/commit/118a70c3cd4c6104a60db58ce6cead34a2e1ae34)) - christoff-linde
- **(.gitignore)** exclude *.work files from ignore - ([3ede7a2](https://github.com/cocogitto/cocogitto/commit/3ede7a2996c185076fd9a9557bcf0f71b57ced1c)) - christoff-linde
- **(api)** initialise package for api logic - ([de77ae0](https://github.com/cocogitto/cocogitto/commit/de77ae00198d7c1c741c41765f1c33306f3194bd)) - christoff-linde
- **(api)** add sample godotenv setup - ([3b8f6e5](https://github.com/cocogitto/cocogitto/commit/3b8f6e5c9853365cf2faf911e1f92eac0e049823)) - christoff-linde
- **(api)** scaffold initial api client & hello-world routes - ([c239f41](https://github.com/cocogitto/cocogitto/commit/c239f4185ba624a7bb9f78c5937fd37b4d1761db)) - christoff-linde
- **(api)** add sensor routes - ([01c00f4](https://github.com/cocogitto/cocogitto/commit/01c00f448a9c5e79a27b7ff05d9213f2bfc5a3fe)) - christoff-linde
- **(api)** add openapi docs generation using huma - ([1acd1ae](https://github.com/cocogitto/cocogitto/commit/1acd1ae23e25d455b872c69efeec7e8228fece90)) - christoff-linde
- **(api)** add sqlc generated code - ([b2f4703](https://github.com/cocogitto/cocogitto/commit/b2f4703328005c9f9aa4f1cbad89f376175bab87)) - christoff-linde
- **(api)** implement GET sensor operations - ([015ded8](https://github.com/cocogitto/cocogitto/commit/015ded8cf5eca8c83163cdcb0922ca1f3f224920)) - christoff-linde
- **(api)** implement CreateSensor & DeleteSensor routes - ([01cc43a](https://github.com/cocogitto/cocogitto/commit/01cc43a8eb7124523cf832b5cbfbc167bcd207f5)) - christoff-linde
- **(api)** add UpdateSensors route & logic - ([45a3597](https://github.com/cocogitto/cocogitto/commit/45a3597c646fcbcf9411f61a9e439d55dc9c8ae7)) - christoff-linde
- **(api)** add sqlc generated code for sensor_metadata - ([c86c8cb](https://github.com/cocogitto/cocogitto/commit/c86c8cb7b683ad8b8a3c2a973646d5da1e6aa45c)) - christoff-linde
- **(api)** implement sensor_metadata routes - ([c3c8744](https://github.com/cocogitto/cocogitto/commit/c3c874415de49c20f878c6befe814bee96135e62)) - christoff-linde
- **(api)** add sqlc generated code for sensor_readings - ([3d1d630](https://github.com/cocogitto/cocogitto/commit/3d1d6300950affc9d64626ceacffb7c483686d51)) - christoff-linde
- **(api)** implement sensor_reading related queries - ([cda9422](https://github.com/cocogitto/cocogitto/commit/cda9422bb6f0f783e2c53bcfd42d3b3f860bf25e)) - christoff-linde
- **(consumer)** initialise package for consumer logic - ([2aeb603](https://github.com/cocogitto/cocogitto/commit/2aeb6033961ea6f8b44eee71467ec196d4a0ea7e)) - christoff-linde
- **(consumer)** add sample godotenv setup - ([c09fcda](https://github.com/cocogitto/cocogitto/commit/c09fcda22e2f410c68ae3428f4a42d46bd0f64cc)) - christoff-linde
- **(consumer)** add dotenv & amqp functionality - ([6bb0787](https://github.com/cocogitto/cocogitto/commit/6bb0787ac523eae7c588b56b42e686f40bb3185a)) - christoff-linde
- **(consumer)** generate database logic with sqlc - ([8c910ec](https://github.com/cocogitto/cocogitto/commit/8c910ec1d14e90cc6712112f72d5c041f21d7daf)) - christoff-linde
- **(consumer)** add appConfig & initial createSensor & createSensorMetadata - ([925e7c9](https://github.com/cocogitto/cocogitto/commit/925e7c9e0f7d058378dcbc84fd1981122a9c1f74)) - christoff-linde
- **(consumer)** update sqlc generated code for sensor & sensor_metadata - ([e022c1f](https://github.com/cocogitto/cocogitto/commit/e022c1f1c55ea129ada4fa866c19e2c38ac67717)) - christoff-linde
- **(consumer)** add logic to create/get sensor & create sensorMetadata - ([4af4233](https://github.com/cocogitto/cocogitto/commit/4af42336b712b1c0371002a42bbeb963dff1815c)) - christoff-linde
- **(consumer)** add logic to write sensor readings into db - ([cad9cfc](https://github.com/cocogitto/cocogitto/commit/cad9cfc783e1cbe27523c5b5155c39866d37a6f6)) - christoff-linde
- **(consumer)** add getSensorByName query - ([3c14dd9](https://github.com/cocogitto/cocogitto/commit/3c14dd93d913b3fccfe210346a15b6684cb54a5f)) - christoff-linde
- **(db)** initialise package for db logic - ([f67fe63](https://github.com/cocogitto/cocogitto/commit/f67fe632d11184a1c8f6551888b18231fe86c600)) - christoff-linde
- **(db)** setup sqlc configs for both applications - ([2b8247f](https://github.com/cocogitto/cocogitto/commit/2b8247fc47db1b3e892680b0d91a1b4a274b825a)) - christoff-linde
- **(db)** add migrations for sensors, sensor_metadata, and sensor_readings - ([d8004af](https://github.com/cocogitto/cocogitto/commit/d8004af0c00dba6b802c06cf765a8d8300574ff4)) - christoff-linde
- **(db)** add sensor sql create query - ([93e5084](https://github.com/cocogitto/cocogitto/commit/93e50845d6cb6df906ae2fc6f61813a2989ea9be)) - christoff-linde
- **(db)** add sensor_readings sql create query - ([d894efc](https://github.com/cocogitto/cocogitto/commit/d894efc93e58520dab5654678a01b2229f1f450d)) - christoff-linde
- **(db)** add sensor_metadata create query - ([226bf28](https://github.com/cocogitto/cocogitto/commit/226bf28b47f0c9e7a4b1cb266388c9b9e2323aff)) - christoff-linde
- **(db)** add GetSensorMetadataForSensorID query - ([279dc8d](https://github.com/cocogitto/cocogitto/commit/279dc8de04603b721081376f4f23db7a04d59683)) - christoff-linde
- **(db)** add GetSensorById query - ([412866d](https://github.com/cocogitto/cocogitto/commit/412866dbd0875ec93a6a34d89d0404937abc20e5)) - christoff-linde
- **(db)** add GetSensorByName query - ([c513902](https://github.com/cocogitto/cocogitto/commit/c513902e8dcba11f302f446bf27966a8b2f8985d)) - christoff-linde
- **(db)** add sensor_readings_daily materialized view - ([4aa0dff](https://github.com/cocogitto/cocogitto/commit/4aa0dffd31315f337ec4ac2cba2f38ebdfaf320e)) - christoff-linde
- **(db)** add sensor_readings_hourly materialized view - ([772f359](https://github.com/cocogitto/cocogitto/commit/772f359d4fa1ccaeae1480532830d47671eec87d)) - christoff-linde
- **(db)** add aggregate policies for hourly & daily materialized views - ([85969a7](https://github.com/cocogitto/cocogitto/commit/85969a7b249209b1447f88103177ccec0dafd26d)) - christoff-linde
- **(db)** add sensor_readings_minute materialized view - ([286682b](https://github.com/cocogitto/cocogitto/commit/286682b71c2ae5139304682ac30393f143eacf3d)) - christoff-linde
- **(db)** add aggregate policy for minute materialized view - ([8a6746e](https://github.com/cocogitto/cocogitto/commit/8a6746e8e3e25d983847368d5ee732a4a7011605)) - christoff-linde
- **(db)** add sensor sqlc queries - ([c49ff44](https://github.com/cocogitto/cocogitto/commit/c49ff44d2e23b8ca73c6b59df8b092e4aff211e4)) - christoff-linde
- **(db)** add sqlc queries for sensor_metadata - ([5dd0d7b](https://github.com/cocogitto/cocogitto/commit/5dd0d7bfd34cc35a0d3fec7595f56eb34a0baaa9)) - christoff-linde
- **(db)** add sqlc queries for sensor_readings - ([f098eef](https://github.com/cocogitto/cocogitto/commit/f098eefc6ee7c37ac115ba540bee3b3e98481a8a)) - christoff-linde
- **(db)** allow null pressure column in sensor_readings - ([b814764](https://github.com/cocogitto/cocogitto/commit/b8147644f20b44a4a578981b3a5b9ffdd7e13569)) - christoff-linde
- **(db)** add getSensorReading query - ([6eb3103](https://github.com/cocogitto/cocogitto/commit/6eb31030db4c72fbe96a27274c293a1d1ce7caa5)) - christoff-linde
- **(go.work)** enable go workspaces for repo - ([28dda46](https://github.com/cocogitto/cocogitto/commit/28dda46952c2989210a41637811c8931ad94f158)) - christoff-linde
- **(justfile)** add commonly used project commands to justfile - ([0dc8d8a](https://github.com/cocogitto/cocogitto/commit/0dc8d8aa53a1ae0e87a39a55d8d660075ac14cb7)) - christoff-linde
- **(justfile)** add default command, goland-related commands, and fmt file - ([dd29afb](https://github.com/cocogitto/cocogitto/commit/dd29afb48a03e01054606df32d00176ebfb17928)) - christoff-linde
- **(vendor)** update packages with dotenv and amqp deps for consumer mod - ([ea0839c](https://github.com/cocogitto/cocogitto/commit/ea0839c7fc2a90a3bdda6f0534a8a8ebc7f8bf22)) - christoff-linde
- **(yaak)** add yaak workspace export file - ([255334c](https://github.com/cocogitto/cocogitto/commit/255334c1a3f9c3881fbd3b9372fa9dbaf035b600)) - christoff-linde
- add .env.example file - ([7d8e001](https://github.com/cocogitto/cocogitto/commit/7d8e001cc24a6526ca7aab8b1e305125ed91853a)) - christoff-linde

### Miscellaneous Chores

- **(.gitignore)** exclude .env.example file - ([e314bb3](https://github.com/cocogitto/cocogitto/commit/e314bb306cb351216a0e6ff469d11337af7023e3)) - christoff-linde
- **(.gitignore)** add bin directory - ([73550d1](https://github.com/cocogitto/cocogitto/commit/73550d117a6006dc4509a9c2fcb4d0a5a67b2bf5)) - christoff-linde
- **(.gitignore)** add tmp/ - ([bbc29cb](https://github.com/cocogitto/cocogitto/commit/bbc29cbf2e6b2ae10d8a82a216107de99413aefd)) - christoff-linde
- **(consumer)** go mod tidy - ([c2cff5b](https://github.com/cocogitto/cocogitto/commit/c2cff5b58c4488f58cf93633366e1fa0480ce981)) - christoff-linde
- **(db)** go mod tidy - ([62c05ad](https://github.com/cocogitto/cocogitto/commit/62c05ad5f87ef9bbb48e474bae68e7e4d7f1d30d)) - christoff-linde
- **(db)** add TODO for converting optional params to pointer types - ([c9f2f2c](https://github.com/cocogitto/cocogitto/commit/c9f2f2c519c475c67877ee76a4e1e03af23f868c)) - christoff-linde
- **(go.work.sum)** update workspace dependencies - ([797682c](https://github.com/cocogitto/cocogitto/commit/797682cd59581c5e650801580dac8cf566ae6cd1)) - christoff-linde
- **(vendor)** go mod tidy - ([49ba2fb](https://github.com/cocogitto/cocogitto/commit/49ba2fb2aed7551992e7d96b5ab1ded47a590f65)) - christoff-linde
- **(vendor)** update explicit deps - ([935bead](https://github.com/cocogitto/cocogitto/commit/935bead9adea886d7da12fd1cd7946e6b8706675)) - christoff-linde
- **(vendor)** update explicit deps with api service deps - ([eef9db2](https://github.com/cocogitto/cocogitto/commit/eef9db29a078b54cf9330ebf3853a5fff1cd55aa)) - christoff-linde
- **(vendor)** update deps for workspace - ([5e17a99](https://github.com/cocogitto/cocogitto/commit/5e17a99afc1f8c94bbfc680a34f2b9998d51d8c4)) - christoff-linde
- **(vendor)** run go work vendor - ([9e1fded](https://github.com/cocogitto/cocogitto/commit/9e1fded9a94db554a7ba681544a7d186a2ce5d6c)) - christoff-linde
- **(yaak)** update json api schema - ([5b654da](https://github.com/cocogitto/cocogitto/commit/5b654dafd9e194d894033147a7f88205d3324363)) - christoff-linde
- vendor code - ([06bd749](https://github.com/cocogitto/cocogitto/commit/06bd7497c0152af4bc19063f82dde89787a7a101)) - christoff-linde
- add todos - ([682663c](https://github.com/cocogitto/cocogitto/commit/682663c3d86a01bfc4e0a4e6741a4de7afaf5c58)) - christoff-linde

### Refactoring

- **(ab)** remove unused query file - ([f606d83](https://github.com/cocogitto/cocogitto/commit/f606d8379e7bf1c181e435aa5f55e7155c7dceb0)) - christoff-linde
- **(api)** remove huma & replace with vanilla gin - ([107a136](https://github.com/cocogitto/cocogitto/commit/107a1368e762cee5586254eae28d005cbc8507fe)) - christoff-linde
- **(api)** move sensor routes into separate file - ([b9c70eb](https://github.com/cocogitto/cocogitto/commit/b9c70eb60e6ba443086022edd90411a45e608311)) - christoff-linde
- **(api)** move sensor_metadata routes into separate file - ([1fbfb92](https://github.com/cocogitto/cocogitto/commit/1fbfb92b998b60cbb8a7d38b88f0e36e6e3908c0)) - christoff-linde
- **(consumer)** update generated sensor queries - ([30509a0](https://github.com/cocogitto/cocogitto/commit/30509a0b7eb2dd7821462a6ee5dd512a040291f9)) - christoff-linde
- **(consumer)** update sqlc generated code for sensorMetadata & sensor - ([839e4a3](https://github.com/cocogitto/cocogitto/commit/839e4a3e49ab9013908f3647101a099163632445)) - christoff-linde
- **(consumer)** use default for readingTimestamp - ([4de38a8](https://github.com/cocogitto/cocogitto/commit/4de38a83192421db441c528b878e474008defd34)) - christoff-linde
- **(consumer)** combine consumer logic into single iot queue - ([c02785b](https://github.com/cocogitto/cocogitto/commit/c02785baa7112c71cd1f78612b52e548cf020268)) - christoff-linde
- **(consumer)** update sqlc generated code - ([6d7deed](https://github.com/cocogitto/cocogitto/commit/6d7deed496ffbb6a77fb4117e8e81406a2d5593c)) - christoff-linde
- **(consumer)** write sensorReading as single commit to db - ([6c32a41](https://github.com/cocogitto/cocogitto/commit/6c32a41ab6f161de8556c2dde89bee3c08616035)) - christoff-linde
- **(consumer)** split main.go into smaller separate files - ([adb379a](https://github.com/cocogitto/cocogitto/commit/adb379a2199fe91ee18328f504bfc3ec01afdb56)) - christoff-linde
- **(consumer)** add generated types & clean up files - ([24bbfe5](https://github.com/cocogitto/cocogitto/commit/24bbfe59396075bb7cb75fdcc7a80ca290455490)) - christoff-linde
- **(db)** update CreateSensor query to take less parameters - ([420c539](https://github.com/cocogitto/cocogitto/commit/420c53960a3b418e02da22d7bdb8a119b63c9149)) - christoff-linde
- **(db)** update sensor_metadata create timestamp defaults - ([d87203c](https://github.com/cocogitto/cocogitto/commit/d87203cc4f5798e465070c2c24c3810b97ad18b4)) - christoff-linde
- **(db)** remove sensor_unique_id column from sensors table - ([7711d90](https://github.com/cocogitto/cocogitto/commit/7711d90f429593ac744cce9801ac268a2805b0b5)) - christoff-linde
- **(db)** ensure sensor_metadata id is SERIAL and sensor_id is unique - ([fc50217](https://github.com/cocogitto/cocogitto/commit/fc5021712a11f12440b0d45548952f0eadb0801c)) - christoff-linde
- **(db)** use default for reading_timestamp - ([9fe7eb7](https://github.com/cocogitto/cocogitto/commit/9fe7eb71985902c18781cb2d7803dd94ccd313f4)) - christoff-linde
- **(db)** move sensor_type and sensor_location to sensor_metadata table - ([895ff0f](https://github.com/cocogitto/cocogitto/commit/895ff0f3935cbf1cca76f57749907741de4c178a)) - christoff-linde
- **(db)** add refactored sensor metadata fields - ([e42d386](https://github.com/cocogitto/cocogitto/commit/e42d386e42e1bf27fb422560aaa598a1e3f01a16)) - christoff-linde
- **(vendor)** remove google uuid vendor - ([c4528a1](https://github.com/cocogitto/cocogitto/commit/c4528a1fdba1f605ef63002b4a8426fe0b0a249d)) - christoff-linde

### Deps

- **(api)** run mod tidy - ([1027779](https://github.com/cocogitto/cocogitto/commit/102777988ef8517e8bde30f6dea7147e23e0c6c0)) - christoff-linde

### Reafactor

- **(codeql)** update build mode to manual - ([084f0dd](https://github.com/cocogitto/cocogitto/commit/084f0dda9208541d631c7cdb7e6828fc0368312d)) - christoff-linde

<!-- generated by git-cliff -->
