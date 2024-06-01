import { getServerSession as originalGetServerSession } from "next-auth";
import GoogleProvider from "next-auth/providers/google";
import type { Session, NextAuthOptions } from "next-auth";
import type { JWT } from "next-auth/jwt";

type JWTProps = {
  token: JWT;
  user: Session["user"];
};

type SessionProps = {
  session: Session;
  token: JWT;
};

export const authOptions: NextAuthOptions = {
  session: {
    strategy: "jwt",
  },
  providers: [
    GoogleProvider({
      clientId: process.env.GOOGLE_ID || "",
      clientSecret: process.env.GOOGLE_SECRET || "",
    }),
  ],
  callbacks: {
    // OAuth認証が完了したタイミングでJWTが作成され、jwt関数が実行される。
    async jwt({ token, user }: JWTProps) {
      // TODO: コールバック時に得られる email アドレスを起点に DB からユーザーレコードを参照
      const dbUser = false;

      // TODO: DB にユーザーレコードがなければ、新規作成された Token を session 関数に渡す
      if (!dbUser) {
        if (user) token.id = user.id;
      }
      return token;

      // TODO: DB にユーザーレコードがあれば、保存済み Token として session 関数に渡す
      return {
        // id: dbUser.id,
        // name: dbUser.name,
        // email: dbUser.email,
        // picture: dbUser.image,
      };
    },
    // jwt関数の戻り値はsession関数に渡される。
    async session({ session, token }: SessionProps) {
      if (session.user) {
        session.user.id = token.id as string;
        session.user.name = token.name;
        session.user.email = token.email;
        session.user.image = token.picture;
      }
      // getServerSession が参照する session
      return session;
    },
  },
};

export const getServerSession = async () => {
  console.log("getServerSession実行");
  return originalGetServerSession(authOptions);
};
