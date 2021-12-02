// pages/api/auth/[...nextauth].js
import NextAuth from "next-auth";
import GithubProvider from "next-auth/providers/github";

export default (req, res) =>
    NextAuth(req, res, {
        providers: [
            GithubProvider({
                clientId: process.env.NEXT_PUBLIC_GITHUB_CLIENT_ID,
                clientSecret: process.env.NEXT_PUBLIC_GITHUB_CLIENT_SECRET,
            }),
        ],
        debug: process.env.NODE_ENV === "development",
        secret: process.env.NEXT_PUBLIC_AUTH_SECRET,
        jwt: {
            secret: process.env.NEXT_PUBLIC_JWT_SECRET,
        },
        callbacks: {
            async redirect(url, baseUrl) {
                return "/";
            },
        },
    });