# Metadata


## Top Level Resource Classes

| General              | Domain Model       | JSON API | RDF Mapping         | SWISSUbase mapping   |
| -------------------- | ------------------ | -------- | ------------------- | -------------------- |
| Project              | Project            | object   | :Project            | StudyVersion (300)   |
| Dataset              | Dataset            | object   | :Dataset            | DatasetVersion (400) |
| Person               | Person             | object   | <!-- XXX -->        | <!-- XXX -->         |
| Organization         | Organization       | object   | <!-- XXX -->        | <!-- XXX -->         |
| Address              | Address            | object   | :Address            | Addresses (112/207)  |
| Grant                | Grant              | object   | :Grant              | <!-- XXX -->?        |
| Data Management Plan | DataManagementPlan | object   | :DataManagementPlan | <!-- XXX -->?        |
| License              | License            | object   | :License            | UserContract (800)?  |




## Properties on Resource Classes

### Project

| Property             | Note                                                           | Cardinality                                 | Domain Model                                       | Domain Model Type             | JSON API             | JSON Type                                    | RDF Mapping                           | RDF type                     | RDF Cardinality | SWISSUbase mapping                                                      |
| -------------------- | -------------------------------------------------------------- | ------------------------------------------- | -------------------------------------------------- | ----------------------------- | -------------------- | -------------------------------------------- | ------------------------------------- | ---------------------------- | --------------- | ----------------------------------------------------------------------- |
| ID                   | internal ID                                                    | 1                                           | `id`                                               | UUID                          | `__id`               | string                                       | IRI                                   | IRI                          | -               | -                                                                       |
| type                 | internal type                                                  | 1                                           | `type`                                             | String                        | `__type`             | string                                       | `@type` / `rdf:type`                  | rdf:type                     | -               | -                                                                       |
| created at           | internal creation date                                         | 1                                           | `createdAt`                                        | Date                          | `__createdAt`        | string                                       | <!-- XXX -->                          | <!-- XXX -->                 | -               | -                                                                       |
| created by           | ID of the user who created the resource                        | 1                                           | `createdBy`                                        | User                          | `__createdBy`        | string                                       | <!-- XXX -->                          | <!-- XXX -->                 | -               | -                                                                       |
| modified at          | internal last modification date                                | 0-1                                         | `modifiedAt`                                       | Date                          | `__modifiedAt`       | string                                       | <!-- XXX -->                          | <!-- XXX -->                 | -               | -                                                                       |
| modified by          | ID of the user who last the resource                           | 0-1                                         | `modifiedBy`                                       | User                          | `__modifiedBy`       | string                                       | <!-- XXX -->                          | <!-- XXX -->                 | -               | -                                                                       |
| deleted at           | internal deletion date                                         | 0-1                                         | `deletedAt`                                        | Date                          | `__deletedAt`        | string                                       | <!-- XXX -->                          | <!-- XXX -->                 | -               | -                                                                       |
| deleted by           | ID of the user who deleted the resource                        | 0-1                                         | `deletedBy`                                        | User                          | `__deletedBy`        | string                                       | <!-- XXX -->                          | <!-- XXX -->                 | -               | -                                                                       |
| (†) short code       | deprecated internal short code                                 | 1                                           | `shortcode`                                        | Shortcode                     | `shortcode`          | string                                       | `:hasShortcode`                       | xsd:string                   | 1               | ?                                                                       |
| name                 | project name                                                   | 1                                           | `name`                                             | String                        | `name`               | string                                       | `:hasName`                            | xsd:string                   | 1               | Title (302)                                                             |
| description          | project description                                            | 1                                           | `description`                                      | MultiLanguageText             | `description`        | object (multi-language text)                 | `:hasDescription`                     | xsd:string (with `@en` etc.) | 1-n             | Abstract (314)                                                          |
| start date           | start date of the project                                      | 1                                           | `startDate`                                        | Date                          | `startDate`          | string (`YYYY-MM-DD`)                        | `:hasStartDate`                       | xsd:date                     | 1               | Start date (308)                                                        |
| teaser text          | short teaser text for previewing the project                   | 1                                           | `teaserText`                                       | String                        | `teaserText`         | string                                       | ? <!-- TODO: what to call this? -->   | xsd:string                   | 1               | ?                                                                       |
| datasets             | references to all datasets that are part of the project        | 1-n                                         | `datasets`                                         | Dataset[]                     | `datasets`           | array of string (ID)                         | - <!-- TODO: add? or leave as is? --> | -                            | 1-n             | <!-- TODO: note: they link it in dataset (400.1) -->                    |
| keywords             | keywords describing the project                                | 1-n                                         | `keywords`                                         | MultiLanguageText[]           | `keywords`           | array of object (multi-language text)        | `:hasKeywords`                        | xsd:string                   | 1-n             | ?                                                                       |
| disciplines          | discipline of research the project belongs to                  | 1-n                                         | `disciplineTexts` and `disciplineRefs`             | MultiLanguageText[] and URL[] | `disciplines`        | array of object (multi-language text or URL) | `:hasDiscipline`                      | schema:URL or xsd:string     | 1-n             | ? <!-- TODO: ideally create our own skos on basis of SNF vocabulary --> |
| temporal coverage    | time period covered by the project                             | 1-n                                         | `temporalCoverageTexts` and `temporalCoverageRefs` | MultiLanguageText[] and URL[] | `temporalCoverage`   | array of object (multi-language text or URL) | `:hasTemporalCoverage`                | schema:URL or xsd:string     | 1-n             | Period (311)                                                            |
| funders              | funding persons or organizations                               | 1-n                                         | `funderPersons` and `funderOrganizations`          | Person[] and Organization[]   | `funders`            | array of string (ID)                         | `:hasFunder`                          | :Person or :Organization     | 1-n             | Funding (318)                                                           |
| primary URL          | project's primary URL, pointing to the DaSCH                   | 1                                           | `url`                                              | URL                           | `url`                | object (`URL`)                               | `:hasURL`                             | schema:URL                   | 1               | ? <!-- TODO: is this or the next one 329  -->                           |
| secondary URL        | optional secondary URL, pointing to a specific project website | 0-1                                         | `secondaryURL`                                     | URL                           | `secondaryURL`       | object (`URL`)                               | `:hasSecondaryURL`                    | schema:URL                   | 0-1             | ?                                                                       |
| data management plan | data management plan                                           | 0-1                                         | `dataManagementPlan`                               | DataManagementPlan            | `dataManagementPlan` | string (ID)                                  | `:hasDataManagementPlan`              | :DataManagementPlan          | 0-1             | ?                                                                       |
| end date             | end date of the project                                        | 0-1                                         | `endDate`                                          | Date                          | `endDate`            | string (`YYYY-MM-DD`)                        | `:hasEndDate`                         | xsd:date                     | 0-1             | End date (309)                                                          |
| contact point        | contact person or organization                                 | 0-1                                         | `contactPerson` or `contactOrganization`           | Person or Organization        | `contactPoint`       | string (ID)                                  | `:hasContactPoint`                    | :Person or :Organization     | 0-1             | Authors (306) -> only persons                                           |
| how-to-cite          | representation how to correctly cite the project               | 0-1 <!-- TODO: should this be optional? --> | `howToCite`                                        | String                        | `howToCite`          | string                                       | `:hasHowToCite`                       | xsd:string                   | 0-1             | ?                                                                       |
| publications         | scientific works published in the context of the project       | 0-n                                         | `publications`                                     | String[]                      | `publications`       | array of string                              | `:hasPublication`                     | xsd:string                   | 0-n             | Publications (319)                                                      |
| grants               | financial grants granted to the project                        | 0-n                                         | `grants`                                           | Grant[]                       | `grants`             | array of string (ID)                         | `:hasGrant`                           | :Grant                       | 0-n             | ?                                                                       |
| alternativeNames     | alternative names of the project                               | 0-n                                         | `alternativeNames`                                 | MultiLanguageText[]           | `alternativeNames`   | object (multi-language text)                 | `:hasAlternativeName`                 | xsd:string                   | 0-n             | ?                                                                       |
| ___________________  | _________________________________                              | ___                                         | __________________________                         | ________________________      | ____________________ | ___________________________                  | ___________________________           | ________________             | ___             | ________________                                                        |



### Dataset

| Property            | Note                                                                    | Cardinality | Domain Model                           | Domain Model Type             | JSON API             | JSON Type                                    | RDF Mapping                                                     | RDF type                 | RDF Cardinality | SWISSUbase mapping             |
| ------------------- | ----------------------------------------------------------------------- | ----------- | -------------------------------------- | ----------------------------- | -------------------- | -------------------------------------------- | --------------------------------------------------------------- | ------------------------ | --------------- | ------------------------------ |
| ID                  | internal ID                                                             | 1           | `id`                                   | UUID                          | `__id`               | string                                       | IRI                                                             | IRI                      | -               | -                              |
| type                | internal type                                                           | 1           | `type`                                 | String                        | `__type`             | string                                       | `@type` / `rdf:type`                                            | rdf:type                 | -               | -                              |
| created at          | internal creation date                                                  | 1           | `createdAt`                            | Date                          | `__createdAt`        | string                                       | <!-- XXX -->                                                    | <!-- XXX -->             | -               | -                              |
| created by          | ID of the user who created the resource                                 | 1           | `createdBy`                            | User                          | `__createdBy`        | string                                       | <!-- XXX -->                                                    | <!-- XXX -->             | -               | -                              |
| modified at         | internal last modification date                                         | 0-1         | `modifiedAt`                           | Date                          | `__modifiedAt`       | string                                       | <!-- XXX -->                                                    | <!-- XXX -->             | -               | -                              |
| modified by         | ID of the user who last the resource                                    | 0-1         | `modifiedBy`                           | User                          | `__modifiedBy`       | string                                       | <!-- XXX -->                                                    | <!-- XXX -->             | -               | -                              |
| deleted at          | internal deletion date                                                  | 0-1         | `deletedAt`                            | Date                          | `__deletedAt`        | string                                       | <!-- XXX -->                                                    | <!-- XXX -->             | -               | -                              |
| deleted by          | ID of the user who deleted the resource                                 | 0-1         | `deletedBy`                            | User                          | `__deletedBy`        | string                                       | <!-- XXX -->                                                    | <!-- XXX -->             | -               | -                              |
| title               | title of the dataset                                                    | 1           | `title`                                | String                        | `title`              | string                                       | `:hasTitle`                                                     | xsd:string               | 1               | Title (402)                    |
| accessConditions    | conditions under which the dataset can be accessed                      | 1           | `accessConditions`                     | String                        | `accessConditions`   | string                                       | `:hasConditionsOfAccess` <!-- TODO: adjust name here aswell --> | xsd:string               | 1               | Special permission (413)       |
| how-to-cite         | representation how to correctly cite the dataset                        | 1           | `howToCite`                            | String                        | `howToCite`          | string                                       | `:hasHowToCite`                                                 | xsd:string               | 1               | Bibliographical citation (409) |
| status              | the status of the dataset                                               | 1           | `status`                               | String (enum)                 | `status`             | string (enum)                                | `:hasStatus`                                                    | xsd:string               | 1               | ?                              |
| abstracts           | abstracts describing the dataset                                        | 1-n         | `abstractTexts` and `abstractURLs`     | MultiLanguageText[] and URL[] | `abstract`           | array of object (multi-language text or URL) | `:hasAbstract`                                                  | xsd:string or schema:URL | 1-n             | ?                              |
| types of data       | the types of data contained in the dataset                              | 1-n         | `typesOfData`                          | String[] (enum)               | `typeOfData`         | array of string (enum)                       | `:hasTypeOfData`                                                | xsd:string               | 1-n             | ?                              |
| licenses            | the license(s) under which the data is made available                   | 1-n         | `licenses`                             | License[]                     | `licenses`           | object (license)                             | `:hasLicense`                                                   | <!-- XXX -->             | 1-n             | ?                              |
| languages           | the languages used in the dataset                                       | 1-n         | `languages`                            | MultiLanguageText[]           | `languages`          | object (multi-language text)                 | `:hasLanguage`                                                  | xsd:string               | 1-n             | ?                              |
| attributions        | attributions of contributions to the dataset by people or organizations | 1-n         | `attributions`                         | Attribution[]                 | `attributions`       | object (attribution)                         | `:hasQualifiedAttribution`                                      | prov:Attribution         | 1-n             | ?                              |
| date published      | publication date of the dataset                                         | 0-1         | `datePublished`                        | Date                          | `datePublished`      | string                                       | `:hasDatePublished`                                             | xsd:date                 | 0-1             | ?                              |
| date created        | creation date of the dataset                                            | 0-1         | `dateCreated`                          | Date                          | `dateCreated`        | string                                       | `:hasDateCreated`                                               | xsd:Date                 | 0-1             | ?                              |
| date modified       | last modification date of the dataset                                   | 0-1         | `dateModified`                         | Date                          | `dateModified`       | string                                       | `:hasDateModified`                                              | xsd:Date                 | 0-1             | ?                              |
| distribution        | distribution URL of the entire dataset                                  | 0-1         | `distribution`                         | URL                           | `distribution`       | object (URL)                                 | `:hasDistribution`                                              | schema:DataDownload      | 0-1             | ?                              |
| alternative titles  | alternative titles of the dataset                                       | 0-n         | `alternativeTitles`                    | MultiLanguageText[]           | `alternativeTitles`  | array of object (multi-language text)        | `:hasAlternativeTitle`                                          | xsd:text                 | 0-n             | ?                              |
| URLs                | URLs pointing to representations of this dataset                        | 0-n         | `urls`                                 | URL[]                         | `urls`               | array of  object (URL)                       | `:hasURL` / `:sameAs` <!-- TODO: which one? -->                 | schema:URL               | 0-n             | ?                              |
| additional          | additional information on the dataset                                   | 0-n         | `additionalTexts` and `additionalURLs` | MultiLanguageText[] and URL[] | `additional`         | array of object (multi-language text or URL) | `:hasAdditional`                                                | xsd:string or schema:URL | 0-n             | Additional Information (404)   |
| ___________________ | _______________________________________                                 | ___         | __________________________             | ________________________      | ____________________ | ___________________________                  | ___________________________                                     | ________________         | ___             | ________________               |


<!-- TODO: make access conditions a fixed vocabulary: open, restricted, closed -->


#### Status

Dataset status can have one of the following values:

- `In planning`
- `Ongoing`
- `On hold`
- `Finished`

#### Types of Data

The following Values are allowed in "Types of Data":

- `XML`
- `Text`
- `Image`
- `Video`
- `Audio`


### Address

| Property      | Note                                    | Cardinality | Domain Model  | Domain Model Type | JSON API        | JSON Type | RDF Mapping               | RDF type     | RDF Cardinality | SWISSUbase mapping               |
| ------------- | --------------------------------------- | ----------- | ------------- | ----------------- | --------------- | --------- | ------------------------- | ------------ | --------------- | -------------------------------- |
| ID            | internal ID                             | 1           | `id`          | UUID              | `__id`          | string    | IRI                       | IRI          | -               | -                                |
| type          | internal type                           | 1           | `type`        | String            | `__type`        | string    | `@type` / `rdf:type`      | rdf:type     | -               | -                                |
| created at    | internal creation date                  | 1           | `createdAt`   | Date              | `__createdAt`   | string    | <!-- XXX -->              | <!-- XXX --> | -               | -                                |
| created by    | ID of the user who created the resource | 1           | `createdBy`   | User              | `__createdBy`   | string    | <!-- XXX -->              | <!-- XXX --> | -               | -                                |
| modified at   | internal last modification date         | 0-1         | `modifiedAt`  | Date              | `__modifiedAt`  | string    | <!-- XXX -->              | <!-- XXX --> | -               | -                                |
| modified by   | ID of the user who last the resource    | 0-1         | `modifiedBy`  | User              | `__modifiedBy`  | string    | <!-- XXX -->              | <!-- XXX --> | -               | -                                |
| deleted at    | internal deletion date                  | 0-1         | `deletedAt`   | Date              | `__deletedAt`   | string    | <!-- XXX -->              | <!-- XXX --> | -               | -                                |
| deleted by    | ID of the user who deleted the resource | 0-1         | `deletedBy`   | User              | `__deletedBy`   | string    | <!-- XXX -->              | <!-- XXX --> | -               | -                                |
| street        | street                                  | 1           | `street`      | String            | `street`        | string    | `schema:streetAddress`    | xsd:string   | 1               | Address line 1 (112.1 / 207.1)   |
| postal code   | postal code of the locality             | 1           | `postalCode`  | String            | `postalCode`    | string    | `schema:postalCode`       | xsd:string   | 1               | Zip code (112.4 / 207.4)         |
| locality      | the name of the locality                | 1           | `locality`    | String            | `locality`      | string    | `schema:addressLocality`  | xsd:string   | 1               | City (112.5 / 207.5)             |
| country       | country                                 | 1           | `country`     | String            | `country`       | string    | `schema:addressCountry`   | xsd:string   | 1               | Country (112.7 / 207.7)          |
| canton        | canton/state within the country         | 0-1         | `canton`      | String            | `canton`        | string    | `schema:addressRegion`    | xsd:string   | 0-1             | Canton (112.6 / 207.6)           |
| additional    | additional information (post box, c/o)  | 0-1         | `additional`  | String            | `additional`    | string    | ? <!-- XXX: ? -->         | xsd:string   | 0-1             | 112.2? 112.3? <!-- XXX: ? -->    |
| _____________ | _______________________________________ | ___         | _____________ | _________________ | _______________ | ________  | _________________________ | ___________  | ___             | ________________________________ |


### Grant

| Property      | Note                                    | Cardinality | Domain Model                              | Domain Model Type           | JSON API        | JSON Type                    | RDF Mapping               | RDF type                 | RDF Cardinality | SWISSUbase mapping          |
| ------------- | --------------------------------------- | ----------- | ----------------------------------------- | --------------------------- | --------------- | ---------------------------- | ------------------------- | ------------------------ | --------------- | --------------------------- |
| ID            | internal ID                             | 1           | `id`                                      | UUID                        | `__id`          | string                       | IRI                       | IRI                      | -               | -                           |
| type          | internal type                           | 1           | `type`                                    | String                      | `__type`        | string                       | `@type` / `rdf:type`      | rdf:type                 | -               | -                           |
| created at    | internal creation date                  | 1           | `createdAt`                               | Date                        | `__createdAt`   | string                       | <!-- XXX -->              | <!-- XXX -->             | -               | -                           |
| created by    | ID of the user who created the resource | 1           | `createdBy`                               | User                        | `__createdBy`   | string                       | <!-- XXX -->              | <!-- XXX -->             | -               | -                           |
| modified at   | internal last modification date         | 0-1         | `modifiedAt`                              | Date                        | `__modifiedAt`  | string                       | <!-- XXX -->              | <!-- XXX -->             | -               | -                           |
| modified by   | ID of the user who last the resource    | 0-1         | `modifiedBy`                              | User                        | `__modifiedBy`  | string                       | <!-- XXX -->              | <!-- XXX -->             | -               | -                           |
| deleted at    | internal deletion date                  | 0-1         | `deletedAt`                               | Date                        | `__deletedAt`   | string                       | <!-- XXX -->              | <!-- XXX -->             | -               | -                           |
| deleted by    | ID of the user who deleted the resource | 0-1         | `deletedBy`                               | User                        | `__deletedBy`   | string                       | <!-- XXX -->              | <!-- XXX -->             | -               | -                           |
| funders       | funding persons or organizations        | 1-n         | `funderPersons` and `funderOrganizations` | Person[] and Organization[] | `funders`       | array of string (ID)         | `:hasFunder`              | :Person or :Organization | 1-n             | ?                           |
| number        | the official grant number               | 0-1         | `number`                                  | String                      | `number`        | string                       | `:hasNumber`              | xsd:string               | 0-1             | ?                           |
| name          | name of the grant type                  | 0-1         | `name`                                    | MultiLanguageText           | `name`          | object (multi-language text) | `:hasName`                | xsd:string               | 0-n             | ?                           |
| url           | url of the grant                        | 0-1         | `url`                                     | URL                         | `url`           | object (URL)                 | `:hasURL`                 | schema:URL               | 0-1             | ?                           |
| _____________ | _______________________________________ | ___         | _______________________                   | _________________           | _______________ | ________________             | _________________________ | ___________              | ___             | ___________________________ |


### Data Management Plan

| Property      | Note                                             | Cardinality | Domain Model      | Domain Model Type | JSON API        | JSON Type        | RDF Mapping               | RDF type      | RDF Cardinality | SWISSUbase mapping          |
| ------------- | ------------------------------------------------ | ----------- | ----------------- | ----------------- | --------------- | ---------------- | ------------------------- | ------------- | --------------- | --------------------------- |
| ID            | internal ID                                      | 1           | `id`              | UUID              | `__id`          | string           | IRI                       | IRI           | -               | -                           |
| type          | internal type                                    | 1           | `type`            | String            | `__type`        | string           | `@type` / `rdf:type`      | rdf:type      | -               | -                           |
| created at    | internal creation date                           | 1           | `createdAt`       | Date              | `__createdAt`   | string           | <!-- XXX -->              | <!-- XXX -->  | -               | -                           |
| created by    | ID of the user who created the resource          | 1           | `createdBy`       | User              | `__createdBy`   | string           | <!-- XXX -->              | <!-- XXX -->  | -               | -                           |
| modified at   | internal last modification date                  | 0-1         | `modifiedAt`      | Date              | `__modifiedAt`  | string           | <!-- XXX -->              | <!-- XXX -->  | -               | -                           |
| modified by   | ID of the user who last the resource             | 0-1         | `modifiedBy`      | User              | `__modifiedBy`  | string           | <!-- XXX -->              | <!-- XXX -->  | -               | -                           |
| deleted at    | internal deletion date                           | 0-1         | `deletedAt`       | Date              | `__deletedAt`   | string           | <!-- XXX -->              | <!-- XXX -->  | -               | -                           |
| deleted by    | ID of the user who deleted the resource          | 0-1         | `deletedBy`       | User              | `__deletedBy`   | string           | <!-- XXX -->              | <!-- XXX -->  | -               | -                           |
| available     | indicates if a data management plan is available | 0-1         | `available`       | Boolean           | `available`     | boolean          | `:isAvailable`            | xsd:boolean   | 0-1             | ?                           |
| URL           | URL to the data management plan                  | 0-1         | `url`             | URL               | `url`           | object (URL)     | `hasURL`                  | schema:URL    | 0-1             | ?                           |
| _____________ | _______________________________________          | ___         | _________________ | _________________ | _______________ | ________________ | _________________________ | _____________ | ___             | ___________________________ |


### License

| Property      | Note                                        | Cardinality | Domain Model      | Domain Model Type | JSON API        | JSON Type        | RDF Mapping               | RDF type      | RDF Cardinality | SWISSUbase mapping          |
| ------------- | ------------------------------------------- | ----------- | ----------------- | ----------------- | --------------- | ---------------- | ------------------------- | ------------- | --------------- | --------------------------- |
| ID            | internal ID                                 | 1           | `id`              | UUID              | `__id`          | string           | IRI                       | IRI           | -               | -                           |
| type          | internal type                               | 1           | `type`            | String            | `__type`        | string           | `@type` / `rdf:type`      | rdf:type      | -               | -                           |
| created at    | internal creation date                      | 1           | `createdAt`       | Date              | `__createdAt`   | string           | <!-- XXX -->              | <!-- XXX -->  | -               | -                           |
| created by    | ID of the user who created the resource     | 1           | `createdBy`       | User              | `__createdBy`   | string           | <!-- XXX -->              | <!-- XXX -->  | -               | -                           |
| modified at   | internal last modification date             | 0-1         | `modifiedAt`      | Date              | `__modifiedAt`  | string           | <!-- XXX -->              | <!-- XXX -->  | -               | -                           |
| modified by   | ID of the user who last the resource        | 0-1         | `modifiedBy`      | User              | `__modifiedBy`  | string           | <!-- XXX -->              | <!-- XXX -->  | -               | -                           |
| deleted at    | internal deletion date                      | 0-1         | `deletedAt`       | Date              | `__deletedAt`   | string           | <!-- XXX -->              | <!-- XXX -->  | -               | -                           |
| deleted by    | ID of the user who deleted the resource     | 0-1         | `deletedBy`       | User              | `__deletedBy`   | string           | <!-- XXX -->              | <!-- XXX -->  | -               | -                           |
| license       | URL pointing to the license itself          | 1           | `license`         | URL               | `license`       | object (URL)     | `:hasLicenseURL`          | schema:URL    | 1               | ?                           |
| date          | licensing date                              | 1           | `date`            | Date              | `date`          | string           | `:hasDate`                | xsd:date      | 1               | ?                           |
| details       | additional information (e.g. license scope) | 0-1         | `details`         | String            | `details`       | string           | `:hasDetails`             | xsd:string    | 0-1             | ?                           |
| _____________ | _______________________________________     | ___         | _________________ | _________________ | _______________ | ________________ | _________________________ | _____________ | ___             | ___________________________ |



## Data Types

### Overview

| General             | Domain Model      | JSON API | RDF Mapping                  |
| ------------------- | ----------------- | -------- | ---------------------------- |
| UUID                | UUID              | string   | xsd:string                   |
| string              | String            | string   | xsd:string                   |
| boolean             | Boolean           | boolean  | xsd:boolean                  |
| URL                 | URL               | object   | schema:URL                   |
| date                | Date              | string   | xsd:date                     |
| multi-language text | MultiLanguageText | object   | xsd:string (with `@en` etc.) |


### API Details

#### URL

```jsonc
{
    "__type": "URL",  // internal type: must always be `URL`
    "type": "URL",  // type: either `URL` or a supported authority file
    "url": "https://www.example.com",  // the actual URL
    "text": "17th Century"  // display text for the URL (optional)
}
```

The following values are accepted in `type`:

- URL
- Geonames
- Pleiades
- Skos
- Periodo
- Chronontology
- GND
- VIAF
- Grid


#### Text (Multi-Language)


Representation of the same text in multiple languages. (Minimum 1 language.)

The keys of the object must be a two character ISO language code string. The property can be any string in the according language.

```jsonc
{
  "en": // ISO language code
    "history", // string in said language
  "de": "Geschichte" // the same for another language
}
```

It is recommended to _always_ include english, furthermore any of the official languages of Switzerland can be used. Other languages _may_ be used but are not likely to be prioritized to be displayed in the front end.


#### Attribution

Attribution attributes one or more roles to a person or organization.

```jsonc
{
  "__type": "Attribution",
  "agent": "ID", // ID of a person ort organization
  "roles": [ // the roles of the person/organization
    "PI",
    "Editor"
  ]
}
```



<!-- TODO: add API stuff for all the classes -->





# Old

## Domain Entities

The domain entities (cf. clean architecture) implemented in the backend are structured as follows.

Note that cases, where the data model supposes a two different types, these are represented as two separate arrays for simplicity; these will be concatenated when the field is requested, and returned as one array through the API.

![domain entities](./domain-entities.svg)  
(Click on the image and open `raw` for a reasonable zoom level.)

## API representation

A metadata set can be serialized to JSON. This JSON data must conform to the [json-schema](schema-metadata.json).  
The following example illustrates how the JSON might look:

![json example](./api-example.svg)  
(Click on the image and open `raw` for a reasonable zoom level.)  
(Cf. `example.json` too.)

The JSON representation is "flat", i.e. not nested, so all top-level types are present in the first level of depth of the JSON document tree. All those objects have a unique `@id` property. Wherever this object is referenced further down in the document, this is done so by this ID.  
(NB: JSON schema does not allow for consistency checks of internal references, so the existence of an object with a given ID can not be guaranteed by JSON validation.)

This JSON contains the entirety of a dataset. Endpoints can be implemented that return only part of the data, e.g. only the project, a list of all datasets, a dataset with a certain ID, etc. However, these endpoints should return an exact subset of the data i the full representation of the metadata.

Other serializations of the data are possible. Eventually, a serialization in RDF should bee added.


## Building the UML Diagrams

The UML diagrams are created using `PlantUML` which itself relies on `GraphViz`.

To create the diagrams, run `yarn run uml`. This requires Java and GraphViz to be installed.
