import type { FormEvent } from "react";
import { useState } from "react";
import { Icon } from "@/components/Icon";
import { PhotoDndUploader } from "@/components/PhotoDndUploader";
import { Typography } from "@/components/Typography";
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
    categoryId: "",
  });
  const handleChangeMeta = (state: State) => {
    setState(state);
  };

  const [photoData, setPhotoData] = useState<Blob>();
  const handleChangeFile = (file: Blob) => {
    setPhotoData(file);
  };

  const handleSubmit = async (event: FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    if (!photoData) return;
    try {
      await fetch(`http://localhost:9992/api/photos`, {
        method: "POST",
        headers: { "Content-Type": "dddd" },
        body: JSON.stringify({
          authorId: "testUser1",
          imageUrl: "testImageUrl",
          title,
          categoryId: "category1",
          description,
        }),
      }).then((res) => {
        if (!res.ok) throw new Error();
        return res.json();
      });
    } catch (err) {
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
