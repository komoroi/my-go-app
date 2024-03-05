import type { NextPage } from "next";
import { useQuery } from "@apollo/client";

import { GetUserDocument } from "../../graphql/dist/client";
import { GetUserQuery } from "../../graphql/dist/client";
import { Heading, Container } from '@chakra-ui/react'
import { Header } from "../components/Header";
import Footer from '../components/Footer';
const Home: NextPage = () => {
  const { data } = useQuery<GetUserQuery>(GetUserDocument);
console.log(data);
  return (
    <><Header></Header><Container>
      {/* <Heading as="h1" fontSize="3xl" mt="10">
        こんにちは、Next.js
      </Heading> */}
      <div>
        {data?.users?.map((user) => (
          <div key={user.id}>
            <div>
              <p>id: {user.id}</p>
              {/* 他のユーザー情報の表示もここに追加 */}
            </div>
            <p>あああ</p>
          </div>
        ))}
      </div>
    </Container>
    <Footer/></>
  );
};

export default Home;