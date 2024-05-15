import { postPhotos } from "@/app/services/postPhotos";
import { NextRequest, NextResponse } from "next/server";

export const corsHeaders = {
  "Access-Control-Allow-Origin": "http://localhost:3000",
  "Access-Control-Allow-Methods": "GET, POST, PUT, DELETE, OPTIONS", // OPTONSを追加
  "Access-Control-Allow-Headers": "Content-Type", // 追加
};

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
        headers: {
          "Access-Control-Allow-Origin": "*",
          "Access-Control-Allow-Methods": "GET, POST, PUT, DELETE, OPTIONS",
          "Access-Control-Allow-Headers": "Content-Type, Authorization",
        },
      }
    );
  } catch (err) {
    return Response.json({ message: "Internal Server Error" }, { status: 500 });
  }
}

export async function OPTIONS() {
  return NextResponse.json({}, { headers: corsHeaders });
}
