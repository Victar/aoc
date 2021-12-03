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
		final Map<String, Validator> validationMap = new LinkedHashMap<>();
		int error = 0;
		final List<String> validTickets = new ArrayList<>();
		for (final String input : data) {
			if (input.contains("-")) {
				/// add validation rule
				final String[] split = StringUtils.split(input, ":");
				final String[] splitNum = StringUtils.split(StringUtils.split(input, ":")[1].trim(), " -or");
				validationMap.putIfAbsent(split[0],
						new Validator(Integer.parseInt(splitNum[0]), Integer.parseInt(splitNum[1]), Integer.parseInt(splitNum[2]),
								Integer.parseInt(splitNum[3])));
			} else if (input.contains(",")) {
				final String[] split = StringUtils.split(input, ",");
				boolean isValid = true;
				for (int i = 0; i < split.length; i++) {
					final int num = Integer.parseInt(split[i]);
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
		final Map<String, Validator> validationMap = new LinkedHashMap<>();
		final Map<Integer, List<Integer>> validPosition = new LinkedHashMap<>();

		for (final String input : data) {
			if (input.contains("-")) {
				/// add validation rule
				final String[] split = StringUtils.split(input, ":");
				final String[] splitNum = StringUtils.split(StringUtils.split(input, ":")[1].trim(), " -or");
				validationMap.putIfAbsent(split[0],
						new Validator(Integer.parseInt(splitNum[0]), Integer.parseInt(splitNum[1]), Integer.parseInt(splitNum[2]),
								Integer.parseInt(splitNum[3])));
			} else if (input.contains(",")) {
				final String[] split = StringUtils.split(input, ",");
				boolean isValid = true;
				for (int i = 0; i < split.length; i++) {
					final int num = Integer.parseInt(split[i]);
					if (!isValid(num, validationMap)) {
						isValid = false;
					}
				}
				if (isValid) {
					for (int i = 0; i < split.length; i++) {
						final int num = Integer.parseInt(split[i]);
						validPosition.putIfAbsent(i, new ArrayList<>());
						validPosition.get(i).add(num);
					}
				}
			}
		}
		final Map<Integer, List<String>> positionSeat = new LinkedHashMap<>();

		for (final Map.Entry<Integer, List<Integer>> entry : validPosition.entrySet()) {
			final Integer position = entry.getKey();
			final List<Integer> currentList = entry.getValue();
			System.out.println("position: " + position);
			for (final Map.Entry<String, Validator> validation : validationMap.entrySet()) {
				final Validator currentValidation = validation.getValue();
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
		final int[] myTicket = { 97, 61, 53, 101, 131, 163, 79, 103, 67, 127, 71, 109, 89, 107, 83, 73, 113, 59, 137, 139 };
		long result = 1;
		for (final Map.Entry<Integer, List<String>> entry : positionSeat.entrySet()) {
			final Integer position = entry.getKey();
			final String seatName = entry.getValue().get(0);
			if (seatName.contains("departure")) {
				result *= myTicket[position];
				System.out.println("position: " + position + " my Ticket: " + myTicket[position]);
			}
		}
		System.out.println(result);
	}

	public void normalizeMap(final Map<Integer, List<String>> positionSeatMap, final List<String> normalizedList) {
		String seatNormalize = null;
		for (final Map.Entry<Integer, List<String>> entry : positionSeatMap.entrySet()) {
			final Integer position = entry.getKey();
			final List<String> seatArray = entry.getValue();
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

			for (final Map.Entry<Integer, List<String>> entry : positionSeatMap.entrySet()) {
				final Integer position = entry.getKey();
				final List<String> seatArray = positionSeatMap.get(position);
				if (seatArray.size() > 1) {
					seatArray.remove(seatNormalize);
				}
			}
			normalizeMap(positionSeatMap, normalizedList);
		}
	}

	public boolean isValid(final int num, final Map<String, Validator> validationMap) {
		for (final Map.Entry<String, Validator> entry : validationMap.entrySet()) {
			if (entry.getValue().isValid(num)) {
				return true;
			}
		}
		return false;
	}

	@Data class Validator {

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

		public boolean isValid(final int num) {
			return num >= this.min && num <= this.max || num >= this.min2 && num <= this.max2;
		}

		public boolean isValidList(final List<Integer> list) {
			for (final int num : list) {
				if (!isValid(num)) {
					return false;
				}
			}
			return true;
		}
	}
}

