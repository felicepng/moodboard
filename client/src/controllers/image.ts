import axios from "axios";
import { GenerateImagesReq } from "../models/models";

export const generateImages = async (req: GenerateImagesReq) => {
  const res = await axios.post(`${window.location.protocol}//${window.location.hostname}:8080/api/images`, req)
  return res.data
}