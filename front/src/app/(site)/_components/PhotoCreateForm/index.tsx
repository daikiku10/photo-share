"use client";

import type { FormEvent } from "react";
import { useState } from "react";
import { useRouter } from "next/navigation";
import { Icon } from "@/components/Icon";
import { PhotoDndUploader } from "@/components/PhotoDndUploader";
import { Typography } from "@/components/Typography";
import { uploadPhoto } from "@/lib/s3";
import { PhotoMeta } from "./PhotoMeta";
import styles from "./style.module.css";

type State = {
  title: string;
  description: string;
  categoryId: string;
};

/** 投稿写真アップロード */
export function PhotoUploader({
  onChange,
}: {
  onChange: (file: Blob) => void;
}) {
  return (
    <PhotoDndUploader
      className={styles.photo}
      areaClassName={styles.area}
      dragActiveClassName={styles.dragActive}
      onChange={onChange}
    >
      {(isDragActive) => (
        <>
          <Icon
            type="upload"
            size="large"
            color={isDragActive ? "orange" : "gray"}
          />
          <Typography>ここに写真をドロップしてください</Typography>
        </>
      )}
    </PhotoDndUploader>
  );
}

/** 写真投稿フォーム */
export function PhotoCreateForm() {
  const [{ title, description, categoryId }, setState] = useState<State>({
    title: "",
    description: "",
    categoryId: "category1",
  });
  const handleChangeMeta = (state: State) => {
    setState(state);
  };

  const [photoData, setPhotoData] = useState<Blob>();
  const handleChangeFile = (file: Blob) => {
    setPhotoData(file);
  };

  const router = useRouter();
  const handleSubmit = async (event: FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    if (!photoData) return;

    // imageUrlを受け取る
    const imageUrl = await uploadPhoto({ photoData });

    try {
      const photo = await fetch(`${process.env.API_HOST}/api/photos`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          authorId: "testUser1", // TODO ユーザー情報
          imageUrl,
          title,
          categoryId: categoryId, // TODO カテゴリID
          description,
        }),
      }).then((res) => {
        if (!res.ok) throw new Error();
        return res.json();
      });
      router.refresh();
      router.push(`photos/${photo.id}`);
    } catch (err) {
      console.log(err);
      window.alert("アップロードに失敗しました。");
    }
  };

  return (
    <form className={styles.form} onSubmit={handleSubmit}>
      <PhotoUploader onChange={handleChangeFile} />
      {/* タイトル・説明・カテゴリID */}
      <PhotoMeta onChange={handleChangeMeta} />
    </form>
  );
}
