import type { GetPhotoSuccessResponse } from "@/generate";

/**
 * BFF層に投稿写真取得を依頼する
 * @param id 投稿写真ID
 * @returns API実行結果
 */
export async function getPhoto(
  id: string
): Promise<{ photo: GetPhotoSuccessResponse }> {
  return fetch(`${process.env.API_HOST}/api/photos/${id}`, {
    next: { tags: [`photos/${id}`] },
  })
    .then((res) => {
      if (!res.ok) throw new Error();
      return res.json();
    })
    .catch(() => {
      // TODO: エラーハンドリング
      window.alert("投稿写真の取得に失敗しました。");
    });
}
