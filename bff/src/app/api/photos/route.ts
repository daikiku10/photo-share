import { NextRequest } from "next/server";

export async function POST(req: NextRequest) {
  console.log("POST実行");
  return Response.json({}, { status: 201 });
}
