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

  const handleSubmit = () => {
    console.log("送信");
  };

  return (
    <form className={styles.form} onSubmit={handleSubmit}>
      <PhotoUploader onChange={handleChangeFile} />
      {/* タイトル・説明・カテゴリID */}
      <PhotoMeta />
    </form>
  );
}
