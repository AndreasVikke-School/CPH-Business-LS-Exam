/** @type {import('next').NextConfig} */
module.exports = {
  reactStrictMode: true,
  env: {
    NEXT_PUBLIC_API_IP: process.env.NEXT_PUBLIC_API_IP,
  }
}
