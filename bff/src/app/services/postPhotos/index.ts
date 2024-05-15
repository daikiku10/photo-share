import { PostPhotoRequest, PostPhotoSuccessResponse } from "@/generate/bff";
import { handleFailed, handleSucceed, path } from "..";

export function postPhotos(
  payload: PostPhotoRequest
): Promise<PostPhotoSuccessResponse> {
  return fetch(path(`/photos`), {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(payload),
  })
    .then(handleSucceed)
    .catch(handleFailed);
}
