// document.getElementById("btnHello").addEventListener("click", async () => {
//   const res = await fetch("/api/hello");
//   const data = await res.json();
//   document.getElementById("response").innerText = data.message;
// });

// document.getElementById("btnSend").addEventListener("click", async () => {
//   const text = document.getElementById("inputText").value;
//   const res = await fetch("/api/echo", {
//     method: "POST",
//     headers: { "Content-Type": "application/json" },
//     body: JSON.stringify({ text }),
//   });
//   const data = await res.json();
//   document.getElementById("echoResponse").innerText = JSON.stringify(data, null, 2);
// });

document.getElementById("formDaftar").addEventListener("submit", async (e) => {
  e.preventDefault();
  
  const data = {
    nama: document.getElementById("nama").value,
    email: document.getElementById("email").value,
    matpel: document.getElementById("matpel").value,
    durasi: document.getElementById("durasi").value,
    jadwal:document.getElementById("jadwal").value,
    noHp: document.getElementById("noHp").value,
  };

  try{
    const res = await fetch("/daftar", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(data)
    });

    const result = await res.json();
    document.getElementById("response").innerText = result.message;
    console.log("✅ Data terkirim:",result)

    alert("✅ Data berhasil dikirim!")
    e.target.reset(); // <-- ini yang mereset form
  }catch(err){
      console.error("❌ Error:",err)
      document.getElementById("response").innerText="Gagal mengirim data"
      alert("❌ Gagal: " + (result.error||res.message || "Terjadi kesalahan"));
  }
});
