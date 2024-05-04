import { useState } from "react";
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
  return <></>;
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
      {/* TODO: PhotoMeta */}
    </form>
  );
}
