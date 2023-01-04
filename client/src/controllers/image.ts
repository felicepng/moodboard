import axios from "axios";
import { GenerateImagesReq } from "../models/models";

export const generateImages = async (search: string) => {
  if (!search) {
    return {
      urls: [],
      prompts: []
    }
  }

  const res = await axios.post(`${window.location.protocol}//${window.location.hostname}:8080/api/images`, { theme: search } as GenerateImagesReq)
  return res.data
}