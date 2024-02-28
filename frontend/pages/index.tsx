import type { NextPage } from "next";
import { useQuery } from "@apollo/client";

import { GetUserDocument } from "../graphql/dist/client";
import { GetUserQuery } from "../graphql/dist/client";

const Home: NextPage = () => {
  const { data } = useQuery<GetUserQuery>(GetUserDocument);
console.log(data);
  return (
    <div style={{ margin: "0 auto", width: "1000px" }}>
      {data?.users?.map((user) => (
        <div key={user.id}>
          <div>
            <p>id: {user.id}</p>
            {/* 他のユーザー情報の表示もここに追加 */}
          </div>
        </div>
      ))}
    </div>
  );
};

export default Home;