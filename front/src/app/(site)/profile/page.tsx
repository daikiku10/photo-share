import { getServerSession } from "@/lib/auth";

type Props = {
  searchParams: { [key: string]: string | string[] | undefined };
};

export default async function Page({}: Props) {
  // ログインユーザーからのリクエストであるか確認
  const session = await getServerSession();
  if (!session || !session.user) {
    // notFound();
  }
  return <>プロフィール画面</>;
}
