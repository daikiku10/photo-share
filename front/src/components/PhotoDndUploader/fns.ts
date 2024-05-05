import { MAX_UPLOAD_PHOTO_HEIGHT, MAX_UPLOAD_PHOTO_WIDTH } from "@/constants";

export function getImageElementFromFile(file: File): Promise<HTMLImageElement> {
  return new Promise((resolve, reject) => {
    const reader = new FileReader();
    reader.readAsDataURL(file);

    reader.onload = () => {
      const image = new Image();
      image.src = reader.result as string;
      image.onload = () => {
        resolve(image);
      };
    };
    reader.onerror = (error) => reject(error);
  });
}

export function resizePhoto({
  image,
}: {
  image: HTMLImageElement;
}): Promise<Blob> {
  return new Promise((resolve) => {
    const canvas = document.createElement("canvas");
    canvas.width = MAX_UPLOAD_PHOTO_WIDTH;
    canvas.height = MAX_UPLOAD_PHOTO_HEIGHT;
    const ctx = canvas.getContext("2d");
    if (!ctx) throw new Error("Failed to get canvas context");
    ctx.drawImage(
      image,
      0,
      0,
      image.width,
      image.height,
      0,
      0,
      MAX_UPLOAD_PHOTO_WIDTH,
      MAX_UPLOAD_PHOTO_HEIGHT
    );
    ctx.canvas.toBlob(
      (result) => {
        if (!result) throw new Error("Failed to convert canvas to blob");
        resolve(result);
      },
      "image/jpeg",
      0.8
    );
  });
}
