export interface Category {
  id: number;
  isOpen: boolean;
  name: string;
  sub: string[];
}
export interface PaginationData {
  currentPage: number;
  currentResultsRange: number[];
  totalCount: number;
  totalPages: number;
}
export interface ProjectMetadata {
  description: string;
  id: string;
  name: string;
  metadata: any[];
}

// TODO: add types 
export interface Metadata {
  project: Project;
  datasets: any[];
  persons: any[];
  organizations: any[];
  grants: any[];
  dataManagementPlan: any;
}

export interface Project {
  __type: "Project";
  __id: string;
  __created: string;
  __modified: string;
  shortcode: string;
  name: string;
  description: Text;
  startDate: string;
  keywords: Text[];
  disciplines: (URL | Text)[];
  temporalCoverage: (URL | Text)[];
  spatialCoverage: URL[];
  urls: URL[];
  funders: string[];
  dataManagementPlan: string;
  endDate: string;
  datasets: string[];
  publications: string[];
  grants: string[];
  alternativeNames: Text[];
  contactPoint: string;
  howToCite: string;
}

export interface Dataset {
  __type: "Dataset";
  __id: string;
  __created: string;
  __modified: string;
  title: string;
  accessConditions: string;
  howToCite: string;
  status: string;
  abstracts: (Text | URL)[];
  typeOfData: string[];
  licenses: URL[];
  languages: Text[];
  attributions: Attribution[];
  alternativeTitles?: Text[];
  datePublished?: string;
  dateCreated?: string;
  dateModified?: string;
  distribution?: URL;
  urls?: URL[];
  documentations?: (Text | URL)[];
}

// export interface Documentations {
//   __type: "Documentations";
//   urls?: URL[];
//   texts?: Text[];
// }

// export interface Abstracts {
//   __type: "Abstracts";
//   urls?: URL[];
//   texts?: Text[];
// }

export interface Attribution {
  __type: "Attribution";
  person: string;
  roles: string[];
}

export interface Address {
  __type: "Address";
  street: string;
  additional?: string;
  postalCode: string;
  locality: string;
  country: string;
}

export interface Person {
  __type: "Person";
  id: string;
  type: string;
  created: string;
  modified: string;
  jobTitles: string[];
  givenNames: string[];
  familyNames: string[];
  affiliation: string[];
  address?: Address;
  emails?: string[];
  authorityRefs?: URL[];
}

export interface URL {
  __type: "URL";
  text?: string;
  type: string;
  url: string;
}

export interface Text {
  // TODO: add to data! is missing so far
  __type: "Text";
  [lang: string]: string
}

export interface Organization {
  __type: "Organization";
  id: string;
  type: string;
  created: string;
  modified: string;
  name: string;
  alternativeNames?: Text[];
  url?: URL;
  email?: string;
  address?: Address;
  authorityRefs?: URL[];
}

export interface Grant {
  __type: "Grant";
  id: string;
  type: string;
  created: string;
  modified: string;
  funders: string[];
  number?: string;
  name?: string;
  url?: URL;
}

export interface DataManagementPlan {
  __type: "DataManagementPlan";
  id: string;
  type: string;
  created: string;
  modified: string;
  available?: boolean;
  url?: URL;
}
