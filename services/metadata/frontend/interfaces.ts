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

export interface Metadata {
  project: Project;
  datasets: any[];
  persons: any[];
  organizations: any[];
  grants: any[];
  dataManagementPlan: any;
}

// TODO: mark optionals as such
export interface Project {
  type: string;
  id: string;
  created: string;
  modified: string;
  shortcode: string;
  name: string;
  description: Text;
  startDate: string;
  keywords: Text[];
  disciplines: (URL | Text)[];
  temporalCoverage: (URL | Text)[];
  spatialCoverage: URL[];
  urls: URL[];
  dataManagementPlan: string;
  endDate: string;
  datasets: string[];
  publications: string[];
  grants: string[];
  alternativeNames: Text[];
  funders: string[];
  contactPoint: string;
  howToCite: string;
}

export interface URL {
  text: string;
  type: string;
  url: string;
}

// export type Text = Map<string, string>

export interface Text {
  [lang: string]: string
}

// export type Text = Record<string, string>

