import { NotesPageContent } from "@/components/public/notes/notes-page-content";

export const metadata = {
  title: "Lab Notes | OPETEER",
  description: "Technical findings, systems observations, and notes by Gerardo M Ardianta (@opeteer).",
};

export default function NotesPage() {
  return (
    <div className="pt-24">
      <NotesPageContent />
    </div>
  );
}
