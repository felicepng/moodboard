import axios from "axios";
import { GenerateImagesReq } from "../models/models";

export const generateImages = async (search: string) => {
  if (!search) {
    return {
      urls: [],
      prompts: []
    }
  }

  try {
    const res = await axios.post(`${window.location.protocol}//${window.location.hostname}:8080/api/images`, { theme: search } as GenerateImagesReq)
    return res.data
  } catch (e) {
    return {
      status: e.response.status,
      message: e.response.data.message
    }
  }
}