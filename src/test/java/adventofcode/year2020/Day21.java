package adventofcode.year2020;

import java.util.ArrayList;
import java.util.HashSet;
import java.util.LinkedHashMap;
import java.util.Map;
import java.util.Set;

import org.junit.Test;

import adventofcode.BaseTest;

import static java.util.stream.Collectors.joining;

public class Day21 extends BaseTest {

	private static final String FILE_NAME = "year2020/day21/input.txt";

	@Test
	public void runBoth() throws Exception {

		final ArrayList<String> data = readStringFromFile(FILE_NAME);

		Map<Set<String>, Set<String>> ingredientsAllergensMap = new LinkedHashMap<>();
		Map<String, Set<String>> allergenIngredients = new LinkedHashMap<>();
		Map<String, Set<String>> ingredientsAllergens = new LinkedHashMap<>();

		for (String input : data) {
			Set<String> ingredientList = new HashSet<>();
			Set<String> allergenList = new HashSet<>();
			String[] inputArr = input.split("\\(contains");
			String[] ingredientArr = inputArr[0].split(" ");
			String[] allergenArr = inputArr[1].substring(0, inputArr[1].length() - 1).split(",");
			for (int i = 0; i < ingredientArr.length; i++) {
				String currentIngredient = ingredientArr[i].trim();
				ingredientList.add(currentIngredient);
				Set<String> currentIngridientAllergens = null;
				if (ingredientsAllergens.get(currentIngredient) != null) {
					currentIngridientAllergens = ingredientsAllergens.get(currentIngredient);
				} else {
					currentIngridientAllergens = new HashSet<>();
					ingredientsAllergens.put(currentIngredient, currentIngridientAllergens);
				}
				for (int j = 0; j < allergenArr.length; j++) {
					currentIngridientAllergens.add(allergenArr[j].trim());
				}
			}
			for (int i = 0; i < allergenArr.length; i++) {
				String currentAllergen = allergenArr[i].trim();
				allergenList.add(allergenArr[i].trim());
				Set<String> currentAllergenIngridients = null;
				if (allergenIngredients.get(currentAllergen) != null) {
					currentAllergenIngridients = allergenIngredients.get(currentAllergen);
				} else {
					currentAllergenIngridients = new HashSet<>();
					allergenIngredients.put(currentAllergen, currentAllergenIngridients);
				}
				for (int j = 0; j < ingredientArr.length; j++) {
					currentAllergenIngridients.add(ingredientArr[j].trim());
				}
			}
			ingredientsAllergensMap.put(ingredientList, allergenList);
		}

		for (Map.Entry<Set<String>, Set<String>> entry : ingredientsAllergensMap.entrySet()) {
			for (String allergen : entry.getValue()) {
				final Set<String> toRemove = new HashSet<String>();
				for (String ingredient : allergenIngredients.get(allergen)) {
					if (!entry.getKey().contains(ingredient)) {
						toRemove.add(ingredient);
					}
				}
				allergenIngredients.get(allergen).removeAll(toRemove);
			}
		}

		int count = 0;
		for (Map.Entry<Set<String>, Set<String>> entry : ingredientsAllergensMap.entrySet()) {
			for (String allergen : entry.getKey()) {
				if (!anyContains(allergen, allergenIngredients)) {
					count++;
				}
			}
		}
		System.out.println("Silver:");
		System.out.println(count);

		//part 2
		Map<String, String> ingredientAllergen = new LinkedHashMap<>();
		for (int i = 0; i < allergenIngredients.size(); i++) {
			for (Map.Entry<String, Set<String>> entry : allergenIngredients.entrySet()) {
				if (entry.getValue().size() == 1) {
					final String value = entry.getValue().iterator().next();
					ingredientAllergen.put(entry.getKey(), value);
					for (Map.Entry<String, Set<String>> entry2 : allergenIngredients.entrySet()) {
						entry2.getValue().remove(value);
					}
				}
			}
		}
		System.out.println("Gold:");
		System.out.println(
				ingredientAllergen.entrySet().stream().sorted(Map.Entry.comparingByKey()).map(Map.Entry::getValue).collect(joining(",")));
	}

	public boolean anyContains(String ingredient, Map<String, Set<String>> allergenIngredients) {
		for (Map.Entry<String, Set<String>> entry : allergenIngredients.entrySet()) {
			if (entry.getValue().contains(ingredient)) {
				return true;
			}
		}
		return false;
	}

}

