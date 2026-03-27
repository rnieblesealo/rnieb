export interface GetDrawingsResponse {
  message: string,
  data: Drawing[]
}

export interface Drawing {
  id: string,
  name: string,
  description: string,
  path: string
}
