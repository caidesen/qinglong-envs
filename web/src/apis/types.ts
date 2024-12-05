/* eslint-disable */
// @ts-ignore

export type CreatePanelParams = {
  clientId: string;
  clientSecret: string;
  name: string;
  url: string;
};

export type deletePanelsIdParams = {
  /** 面板ID */
  id: number;
};

export type getPanelsIdParams = {
  /** 面板ID */
  id: number;
};

export type HTTPError = {
  detail?: unknown;
  message?: string;
};

export type Panel = {
  clientId?: string;
  clientSecret?: string;
  createdAt?: string;
  id?: string;
  name?: string;
  updatedAt?: string;
  url?: string;
};

export type UpdatePanelParams = {
  clientId: string;
  clientSecret: string;
  id: number;
  name: string;
  url: string;
};
