function calculate() {
  const weight = parseFloat(document.getElementById("weight").value);
  const activity = document.getElementById("activity").value;

  if (isNaN(weight) || weight <= 0) {
    alert(getTranslation("invalid_weight"));
    return;
  }

  const proteinPerKg = getProteinPerKg(activity);

  const totalProtein = proteinPerKg * weight;

  const creatine = 0.05 * weight;

  const wheyProtein = (totalProtein * 0.3).toFixed(2);

  const water = (weight * 0.033).toFixed(2);

  insertResults(totalProtein, creatine, wheyProtein, water);
}

function getProteinPerKg(activity) {
  switch (activity) {
    case "sedentary":
      return 0.8;
    case "light":
      return 1.0;
    case "moderate":
      return 1.2;
    case "active":
      return 1.5;
    case "very_active":
      return 1.8;
    default:
      return 1.0;
  }
}

function insertResults(protein, creatine, wheyProtein, water) {
  document.getElementById("protein").innerText = protein.toFixed(2);
  document.getElementById("creatine").innerText = creatine.toFixed(2);
  document.getElementById("whey").innerText = wheyProtein;
  document.getElementById("water").innerText = water;
  document.getElementById("results").classList.remove("hidden");
}

function setLanguage(lang) {
  window.location.href = "/set-lang?lang=" + lang;
}
