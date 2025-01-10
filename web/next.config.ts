import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  // Output in standalone mode for containerized app
  output: "standalone",
  rewrites: async () => {
    return [
      {
        source: "/api/:path*",
        destination: "http://localhost:8080/:path*",
      },
    ];
  },
};

export default nextConfig;
