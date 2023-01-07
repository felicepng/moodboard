export interface GenerateImagesReq {
  theme: string
}

export interface ImageObj {
  urls: string[]
  prompts: string[]
  status?: number
  message?: string
}