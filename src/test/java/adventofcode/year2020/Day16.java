package adventofcode.year2020;

import java.util.ArrayList;
import java.util.LinkedHashMap;
import java.util.List;
import java.util.Map;

import org.apache.commons.lang3.StringUtils;
import org.junit.Test;

import adventofcode.BaseTest;
import lombok.Data;

public class Day16 extends BaseTest {

	@Test public void singleCheck() {

	}

	@Test public void runSilver() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2020/day16/input_sample.txt");
		Map<String, Validator> validationMap = new LinkedHashMap<>();
		int error = 0;
		List<String> validTickets = new ArrayList<>();
		for (String input : data) {
			if (input.contains("-")) {
				/// add validation rule
				String[] split = StringUtils.split(input, ":");
				String[] splitNum = StringUtils.split(StringUtils.split(input, ":")[1].trim(), " -or");
				validationMap.putIfAbsent(split[0],
						new Validator(Integer.parseInt(splitNum[0]), Integer.parseInt(splitNum[1]), Integer.parseInt(splitNum[2]),
								Integer.parseInt(splitNum[3])));
			} else if (input.contains(",")) {
				String[] split = StringUtils.split(input, ",");
				boolean isValid = true;
				for (int i = 0; i < split.length; i++) {
					int num = Integer.parseInt(split[i]);
					if (!isValid(num, validationMap)) {
						error += num;
						isValid = false;
					}
					validTickets.add(input);
				}

			}

		}
		System.out.println(validTickets);
	}

	@Test public void runGold() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2020/day16/input.txt");
		Map<String, Validator> validationMap = new LinkedHashMap<>();
		Map<Integer, List<Integer>> validPosition = new LinkedHashMap<>();

		for (String input : data) {
			if (input.contains("-")) {
				/// add validation rule
				String[] split = StringUtils.split(input, ":");
				String[] splitNum = StringUtils.split(StringUtils.split(input, ":")[1].trim(), " -or");
				validationMap.putIfAbsent(split[0],
						new Validator(Integer.parseInt(splitNum[0]), Integer.parseInt(splitNum[1]), Integer.parseInt(splitNum[2]),
								Integer.parseInt(splitNum[3])));
			} else if (input.contains(",")) {
				String[] split = StringUtils.split(input, ",");
				boolean isValid = true;
				for (int i = 0; i < split.length; i++) {
					int num = Integer.parseInt(split[i]);
					if (!isValid(num, validationMap)) {
						isValid = false;
					}
				}
				if (isValid) {
					for (int i = 0; i < split.length; i++) {
						int num = Integer.parseInt(split[i]);
						validPosition.putIfAbsent(i, new ArrayList<>());
						validPosition.get(i).add(num);
					}
				}
			}
		}
		Map<Integer, List<String>> positionSeat = new LinkedHashMap<>();

		for (Map.Entry<Integer, List<Integer>> entry : validPosition.entrySet()) {
			Integer position = entry.getKey();
			List<Integer> currentList = entry.getValue();
			System.out.println("position: " + position);
			for (Map.Entry<String, Validator> validation : validationMap.entrySet()) {
				Validator currentValidation = validation.getValue();
				if (currentValidation.isValidList(currentList)) {
					System.out.println("   " + validation.getKey());
					positionSeat.putIfAbsent(position, new ArrayList<String>());
					positionSeat.get(position).add(validation.getKey());
				}
			}
		}
		System.out.println(positionSeat);
		normalizeMap(positionSeat, new ArrayList<>());
		System.out.println(positionSeat);
		int[] myTicket = new int[] { 97, 61, 53, 101, 131, 163, 79, 103, 67, 127, 71, 109, 89, 107, 83, 73, 113, 59, 137, 139 };
		long result = 1;
		for (Map.Entry<Integer, List<String>> entry : positionSeat.entrySet()) {
			Integer position = entry.getKey();
			String seatName = entry.getValue().get(0);
			if (seatName.contains("departure")) {
				result *= myTicket[position];
				System.out.println("position: " + position + " my Ticket: " + myTicket[position]);
			}
		}
		System.out.println(result);
	}

	public void normalizeMap(Map<Integer, List<String>> positionSeatMap, List<String> normalizedList) {
		String seatNormalize = null;
		for (Map.Entry<Integer, List<String>> entry : positionSeatMap.entrySet()) {
			Integer position = entry.getKey();
			List<String> seatArray = entry.getValue();
			if (seatArray.size() == 1 && !normalizedList.contains(seatArray.get(0))) {
				System.out.println(position);
				System.out.println(seatArray);
				seatNormalize = seatArray.get(0);
				normalizedList.add(seatNormalize);
				break;
			}
		}
		if (StringUtils.isNotEmpty(seatNormalize)) {
			System.out.println("seatNormalize: " + seatNormalize);

			for (Map.Entry<Integer, List<String>> entry : positionSeatMap.entrySet()) {
				Integer position = entry.getKey();
				List<String> seatArray = positionSeatMap.get(position);
				if (seatArray.size() > 1) {
					seatArray.remove(seatNormalize);
				}
			}
			normalizeMap(positionSeatMap, normalizedList);
		}
	}

	public boolean isValid(int num, Map<String, Validator> validationMap) {
		for (Map.Entry<String, Validator> entry : validationMap.entrySet()) {
			if (entry.getValue().isValid(num)) {
				return true;
			}
		}
		return false;
	}

	@Data
	class Validator {

		int min;
		int max;
		int min2;
		int max2;

		public Validator(final int min, final int max, final int min2, final int max2) {
			this.min = min;
			this.max = max;
			this.min2 = min2;
			this.max2 = max2;

		}

		public boolean isValid(int num) {
			return num >= min && num <= max || num >= min2 && num <= max2;
		}

		public boolean isValidList(List<Integer> list) {
			for (int num : list) {
				if (!isValid(num)) {
					return false;
				}
			}
			return true;
		}
	}
}

