import axios from "axios";
import { GenerateImagesReq, ImageObj } from "../models/models";

export const generateImages = async (arr: string[]) => {
  const theme = arr[0]
  if (!theme) {
    return {
      urls: [],
      prompts: []
    } as ImageObj
  }

  try {
    const res = await axios.post(`${window.location.protocol}//${window.location.hostname}:8080/api/images`, { theme } as GenerateImagesReq)
    return res.data
  } catch (e) {
    return {
      status: e.response.status,
      message: e.response.data.message,
      urls: [],
      prompts: []
    } as ImageObj
  }
}