import { postPhotos } from "@/app/services/postPhotos";
import { NextRequest } from "next/server";

export async function POST(req: NextRequest) {
  // TODO: セッション情報を取得する
  try {
    const body = await req.json();
    const { id } = await postPhotos({
      authorId: "testUser1",
      categoryId: "testCategory1",
      title: body.title,
      description: body.description,
      imageUrl: body.imageUrl,
    });
    return Response.json(
      { id },
      {
        status: 201,
      }
    );
  } catch (err) {
    return Response.json({ message: "Internal Server Error" }, { status: 500 });
  }
}
