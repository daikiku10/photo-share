import { envVars } from "@/constans/enviroments";
import S3 from "aws-sdk/clients/s3";
import { NextRequest } from "next/server";

const s3 = new S3({
  accessKeyId: envVars.AWS_ACCESS_KEY_ID,
  secretAccessKey: envVars.AWS_SECRET_ACCESS_KEY,
  endpoint: envVars.AWS_S3_ENDPOINT,
});

export async function POST(request: NextRequest) {
  // TODO: セッション情報を取得する

  const data = await request.json();

  const presignedPost = await s3.createPresignedPost({
    Bucket: process.env.AWS_S3_PHOTO_BUCKET_NAME,
    Fields: {
      key: data.filename,
      "Content-Type": data.fileType,
    },
    Expires: 60, // seconds
    // Conditions: [["content-length-range", 0]],
  });
  return Response.json({ presignedPost });
}
